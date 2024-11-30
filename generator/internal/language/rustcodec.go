// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package language

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/googleapis/google-cloud-rust/generator/internal/api"
	"github.com/iancoleman/strcase"
)

func validatePackageName(c *rustCodec, newPackage, elementName string) error {
	if c.SourceSpecificationPackageName == newPackage {
		return nil
	}
	// Special exceptions for mixin services
	if newPackage == "google.cloud.location" ||
		newPackage == "google.iam.v1" ||
		newPackage == "google.longrunning" {
		return nil
	}
	return fmt.Errorf("rust codec requires all top-level elements to be in the same package want=%s, got=%s for %s",
		c.SourceSpecificationPackageName, newPackage, elementName)
}

func rustValidate(c *rustCodec, api *api.API) error {
	// Set the source package. We should always take the first service registered
	// as the source package. Services with mixins will register those after the
	// source package.
	if len(api.Services) > 0 {
		c.SourceSpecificationPackageName = api.Services[0].Package
	} else if len(api.Messages) > 0 {
		c.SourceSpecificationPackageName = api.Messages[0].Package
	}
	for _, s := range api.Services {
		if err := validatePackageName(c, s.Package, s.ID); err != nil {
			return err
		}
	}
	for _, s := range api.Messages {
		if err := validatePackageName(c, s.Package, s.ID); err != nil {
			return err
		}
	}
	for _, s := range api.Enums {
		if err := validatePackageName(c, s.Package, s.ID); err != nil {
			return err
		}
	}
	return nil
}

func rustFQEnumName(c *rustCodec, e *api.Enum, _ *api.APIState) string {
	return rustMessageScopeName(c, e.Parent, "") + "::" + rustToPascal(e.Name)
}

func rustFQEnumValueName(c *rustCodec, v *api.EnumValue, state *api.APIState) string {
	return fmt.Sprintf("%s::%s::%s", rustEnumScopeName(c, v.Parent), rustToSnake(v.Parent.Name), rustEnumValueName(v, state))
}

func rustOneOfType(c *rustCodec, o *api.OneOf, _ *api.APIState) string {
	return rustMessageScopeName(c, o.Parent, "") + "::" + rustToPascal(o.Name)
}

// Returns the field type, ignoring any repeated or optional attributes.
func rustBaseFieldType(f *api.Field, state *api.APIState) string {
	if f.Typez == api.MESSAGE_TYPE {
		m, ok := state.MessageByID[f.TypezID]
		if !ok {
			slog.Error("unable to lookup type", "id", f.TypezID)
			return ""
		}
		if m.IsMap {
			key := rustFieldType(m.Fields[0], state)
			val := rustFieldType(m.Fields[1], state)
			return "std::collections::HashMap<" + key + "," + val + ">"
		}
		return rustFQMessageName(m, state)
	} else if f.Typez == api.ENUM_TYPE {
		e, ok := state.EnumByID[f.TypezID]
		if !ok {
			slog.Error("unable to lookup type", "id", f.TypezID)
			return ""
		}
		return rustFQEnumName(e, state)
	} else if f.Typez == api.GROUP_TYPE {
		slog.Error("TODO(#39) - better handling of `oneof` fields")
		return ""
	}
	return scalarFieldType(f)

}

func rustEnumScopeName(c *rustCodec, e *api.Enum) string {
	return rustMessageScopeName(c, e.Parent, "")
}

// Constructor function for RustTemplateData
func NewRustTemplateData(api *api.API, c *rustCodec) *RustTemplateData {
	year, _, _ := time.Now().Date()

	return &RustTemplateData{
		TemplateDir:      rustTemplateDir(c.GenerateModule),
		Name:             api.Name,
		Title:            api.Title,
		Description:      api.Description,
		PackageName:      rustPackageName(api),
		RequiredPackages: rustRequiredPackages(),
		HasServices:      len(api.Services) > 0,
		CopyrightYear:    fmt.Sprintf("%04d", year),
		BoilerPlate: append(licenseHeaderBulk(),
			"",
			" Code generated by sidekick. DO NOT EDIT."),
		Imports: rustImports(),
		DefaultHost: func() string {
			if len(api.Services) > 0 {
				return api.Services[0].DefaultHost
			}
			return ""
		}(),
		Services: mapSlice(api.Services, func(s *api.Service) *RustService {
			return newRustService(s, c, state)
		}),
		Messages: mapSlice(api.Messages, func(m *api.Message) *RustMessage {
			return newRustMessage(m, c, api.State)
		}),
		NameToLower: strings.ToLower(api.Name),
	}
}

// Constructor function for rustMethod
func newRustMethod(m *api.Method, c *rustCodec, state *api.APIState) *RustMethod {
	return &RustMethod{
		BodyAccessor:      rustBodyAccessor(m, state),
		DocLines:          rustFormatDocComments(m.Documentation),
		HTTPMethod:        m.PathInfo.Verb,
		HTTPMethodToLower: strings.ToLower(m.PathInfo.Verb),
		HTTPPathArgs:      rustHTTPPathArgs(m.PathInfo, state),
		HTTPPathFmt:       rustHTTPPathFmt(m.PathInfo, state),
		HasBody:           m.PathInfo.BodyFieldPath != "",
		InputTypeName:     rustMethodInOutTypeName(c, m.InputTypeID, state),
		NameToCamel:       strcase.ToCamel(m.Name),
		NameToPascal:      rustToPascal(m.Name),
		NameToSnake:       strcase.ToSnake(m.Name),
		OutputTypeName:    rustMethodInOutTypeName(c, m.OutputTypeID, state),
		QueryParams: mapSlice(rustQueryParams(m, state), func(s *api.Field) *RustField {
			return newRustField(s, c, state)
		}),
	}
}

func rustMethodInOutTypeName(c *rustCodec, id string, state *api.APIState) string {
	if id == "" {
		return ""
	}
	m, ok := state.MessageByID[id]
	if !ok {
		slog.Error("unable to lookup type", "id", id)
		return ""
	}
	return rustFQMessageName(c, m)
}

func rustFQMessageName(c *rustCodec, m *api.Message) string {
	return rustMessageScopeName(c, m.Parent, m.Package) + "::" + rustToPascal(m.Name)
}

func rustMessageScopeName(c *rustCodec, m *api.Message, childPackageName string) string {
	if m == nil {
		return createRustPackage(sourceSpecificationPackageName, packageMapping, packageName, modulePath)
	}
	if m.Parent == nil {
		return createRustPackage(sourceSpecificationPackageName, packageMapping, m.Package) + "::" + rustToSnake(m.Name, modulePath)
	}
	return rustMessageScopeName(c, m.Parent, m.Package) + "::" + rustToSnake(m.Name)
}

func createRustPackage(sourceSpecificationPackageName string, packageMapping map[string]*rustPackage, packageName, modulePath string) string {
	if packageName == sourceSpecificationPackageName {
		return "crate::" + modulePath
	}
	mapped, ok := packageMapping[packageName]
	if !ok {
		slog.Error("unknown source package name", "name", packageName)
		return packageName
	}
	// TODO(#158) - maybe google.protobuf should not be this special?
	if packageName == "google.protobuf" {
		return mapped.Name
	}
	return mapped.Name + "::model"
}
