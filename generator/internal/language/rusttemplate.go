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
	"strconv"
	"strings"
	"time"

	"github.com/googleapis/google-cloud-rust/generator/internal/api"
	"github.com/googleapis/google-cloud-rust/generator/internal/license"
	"github.com/iancoleman/strcase"
)

type RustTemplateData struct {
	Name              string
	Title             string
	Description       string
	PackageName       string
	PackageVersion    string
	RequiredPackages  []string
	HasServices       bool
	CopyrightYear     string
	BoilerPlate       []string
	Imports           []string
	DefaultHost       string
	Services          []*RustService
	Messages          []*RustMessage
	Enums             []*RustEnum
	NameToLower       string
	NotForPublication bool
	HasFeatures       bool
	Features          []string
}

type RustService struct {
	Methods             []*RustMethod
	NameToSnake         string
	NameToPascal        string
	ServiceNameToPascal string
	NameToCamel         string
	ServiceName         string
	DocLines            []string
	DefaultHost         string
}

type RustMessage struct {
	Fields             []*RustField
	BasicFields        []*RustField
	ExplicitOneOfs     []*RustOneOf
	NestedMessages     []*RustMessage
	Enums              []*RustEnum
	MessageAttributes  []string
	Name               string
	QualifiedName      string
	NameSnakeCase      string
	HasNestedTypes     bool
	DocLines           []string
	IsMap              bool
	IsPageableResponse bool
	PageableItem       *RustField
	ID                 string
	// The FQN is the source specification
	SourceFQN string
	// If true, this is a synthetic message, some generation is skipped for
	// synthetic messages
	HasSyntheticFields bool
}

type RustMethod struct {
	NameToSnake         string
	NameToCamel         string
	NameToPascal        string
	DocLines            []string
	InputTypeName       string
	OutputTypeName      string
	HTTPMethod          string
	HTTPMethodToLower   string
	HTTPPathFmt         string
	HTTPPathArgs        []string
	PathParams          []*RustField
	QueryParams         []*RustField
	HasBody             bool
	BodyAccessor        string
	IsPageable          bool
	ServiceNameToPascal string
	ServiceNameToCamel  string
	ServiceNameToSnake  string
	InputTypeID         string
	InputType           *RustMessage
	OperationInfo       *RustOperationInfo
}

type RustOperationInfo struct {
	MetadataType string
	ResponseType string
}

type RustOneOf struct {
	NameToPascal          string
	NameToSnake           string
	NameToSnakeNoMangling string
	FieldType             string
	DocLines              []string
	Fields                []*RustField
}

type RustField struct {
	NameToSnake           string
	NameToSnakeNoMangling string
	NameToCamel           string
	NameToPascal          string
	DocLines              []string
	FieldAttributes       []string
	FieldType             string
	PrimitiveFieldType    string
	JSONName              string
	AsQueryParameter      string
}

type RustEnum struct {
	Name          string
	NameSnakeCase string
	DocLines      []string
	Values        []*RustEnumValue
}

type RustEnumValue struct {
	DocLines []string
	Name     string
	Number   int32
	EnumType string
}

// newRustTemplateData creates a struct used as input for Mustache templates.
// Fields and methods defined in this struct directly correspond to Mustache
// tags. For example, the Mustache tag {{#Services}} uses the
// [Template.Services] field.
func newRustTemplateData(model *api.API, options map[string]string, outdir string) (*RustTemplateData, []GeneratedFile, error) {
	year, _, _ := time.Now().Date()
	c := &rustCodec{
		generationYear:           fmt.Sprintf("%04d", year),
		modulePath:               "model",
		deserializeWithdDefaults: true,
		extraPackages:            []*rustPackage{},
		packageMapping:           map[string]*rustPackage{},
		version:                  "0.0.0",
	}
	for key, definition := range options {
		switch key {
		case "package-name-override":
			c.packageNameOverride = definition
		case "generate-module":
			value, err := strconv.ParseBool(definition)
			if err != nil {
				return nil, nil, fmt.Errorf("cannot convert `generate-module` value %q to boolean: %w", definition, err)
			}
			c.generateModule = value
		case "module-path":
			c.modulePath = definition
		case "deserialize-with-defaults":
			value, err := strconv.ParseBool(definition)
			if err != nil {
				return nil, nil, fmt.Errorf("cannot convert `deserialize-with-defaults` value %q to boolean: %w", definition, err)
			}
			c.deserializeWithdDefaults = value
		case "copyright-year":
			c.generationYear = definition
		case "not-for-publication":
			value, err := strconv.ParseBool(definition)
			if err != nil {
				return nil, nil, fmt.Errorf("cannot convert `not-for-publication` value %q to boolean: %w", definition, err)
			}
			c.doNotPublish = value
		case "version":
			c.version = definition
		default:
			if !strings.HasPrefix(key, "package:") {
				return nil, nil, fmt.Errorf("unknown Rust codec option %q", key)
			}

			pkgOption, err := parseRustPackageOption(key, definition)
			if err != nil {
				return nil, nil, err
			}
			c.extraPackages = append(c.extraPackages, pkgOption.pkg)
			for _, source := range pkgOption.otherNames {
				c.packageMapping[source] = pkgOption.pkg
			}
		}
	}
	if err := rustValidate(model, c.sourceSpecificationPackageName); err != nil {
		return nil, nil, err
	}
	for _, message := range rustWellKnown {
		model.State.MessageByID[message.ID] = message
	}
	for _, pkg := range c.extraPackages {
		if pkg.requiredByServices {
			pkg.used = rustHasServices(model.State)
		}
	}
	data := &RustTemplateData{
		Name:           model.Name,
		Title:          model.Title,
		Description:    model.Description,
		PackageName:    c.packageName(model),
		PackageVersion: c.packageVersion(),
		HasServices:    len(model.Services) > 0,
		CopyrightYear:  c.copyrightYear(),
		BoilerPlate: append(license.LicenseHeaderBulk(),
			"",
			" Code generated by sidekick. DO NOT EDIT."),
		DefaultHost: func() string {
			if len(model.Services) > 0 {
				return model.Services[0].DefaultHost
			}
			return ""
		}(),
		Services: mapSlice(model.Services, func(s *api.Service) *RustService {
			return newRustService(s, c, model.State)
		}),
		Messages: mapSlice(model.Messages, func(m *api.Message) *RustMessage {
			return newRustMessage(m, c, model.State)
		}),
		Enums: mapSlice(model.Enums, func(e *api.Enum) *RustEnum {
			return newRustEnum(e, c, model.State)
		}),
		NameToLower:       strings.ToLower(model.Name),
		NotForPublication: c.notForPublication(),
	}
	// Delay this until the Codec had a chance to compute what packages are
	// used.
	data.RequiredPackages = c.requiredPackages(outdir)
	c.addStreamingFeature(data, model)

	messagesByID := map[string]*RustMessage{}
	for _, m := range data.Messages {
		messagesByID[m.ID] = m
	}
	for _, s := range data.Services {
		for _, method := range s.Methods {
			if msg, ok := messagesByID[method.InputTypeID]; ok {
				method.InputType = msg
			} else if m, ok := model.State.MessageByID[method.InputTypeID]; ok {
				method.InputType = newRustMessage(m, c, model.State)
			}
		}
	}

	return data, rustGeneratedFiles(c.generateModule, model.State), nil
}

func newRustService(s *api.Service, c *rustCodec, state *api.APIState) *RustService {
	// Some codecs skip some methods.
	methods := filterSlice(s.Methods, func(m *api.Method) bool {
		return c.generateMethod(m)
	})
	return &RustService{
		Methods: mapSlice(methods, func(m *api.Method) *RustMethod {
			return newRustMethod(m, c, state)
		}),
		NameToSnake:         c.toSnake(s.Name),
		NameToPascal:        c.toPascal(s.Name),
		ServiceNameToPascal: c.toPascal(s.Name), // Alias for clarity
		NameToCamel:         c.toCamel(s.Name),
		ServiceName:         s.Name,
		DocLines:            c.formatDocComments(s.Documentation, state),
		DefaultHost:         s.DefaultHost,
	}
}

func newRustMessage(m *api.Message, c *rustCodec, state *api.APIState) *RustMessage {
	hasSyntheticFields := false
	for _, f := range m.Fields {
		if f.Synthetic {
			hasSyntheticFields = true
			break
		}
	}
	return &RustMessage{
		Fields: mapSlice(m.Fields, func(s *api.Field) *RustField {
			return newRustField(s, c, state)
		}),
		BasicFields: func() []*RustField {
			filtered := filterSlice(m.Fields, func(s *api.Field) bool {
				return !s.IsOneOf
			})
			return mapSlice(filtered, func(s *api.Field) *RustField {
				return newRustField(s, c, state)
			})
		}(),
		ExplicitOneOfs: mapSlice(m.OneOfs, func(s *api.OneOf) *RustOneOf {
			return newRustOneOf(s, c, state)
		}),
		NestedMessages: mapSlice(m.Messages, func(s *api.Message) *RustMessage {
			return newRustMessage(s, c, state)
		}),
		Enums: mapSlice(m.Enums, func(s *api.Enum) *RustEnum {
			return newRustEnum(s, c, state)
		}),
		MessageAttributes: c.messageAttributes(m, state),
		Name:              c.messageName(m),
		QualifiedName:     c.fqMessageName(m, state),
		NameSnakeCase:     c.toSnake(m.Name),
		HasNestedTypes: func() bool {
			if len(m.Enums) > 0 || len(m.OneOfs) > 0 {
				return true
			}
			for _, child := range m.Messages {
				if !child.IsMap {
					return true
				}
			}
			return false
		}(),
		DocLines:           c.formatDocComments(m.Documentation, state),
		IsMap:              m.IsMap,
		IsPageableResponse: m.IsPageableResponse,
		PageableItem:       newRustField(m.PageableItem, c, state),
		ID:                 m.ID,
		SourceFQN:          strings.TrimPrefix(m.ID, "."),
		HasSyntheticFields: hasSyntheticFields,
	}
}

func newRustMethod(m *api.Method, c *rustCodec, state *api.APIState) *RustMethod {
	method := &RustMethod{
		BodyAccessor:      c.bodyAccessor(m),
		DocLines:          c.formatDocComments(m.Documentation, state),
		HTTPMethod:        m.PathInfo.Verb,
		HTTPMethodToLower: strings.ToLower(m.PathInfo.Verb),
		HTTPPathArgs:      c.httpPathArgs(m.PathInfo),
		HTTPPathFmt:       c.httpPathFmt(m.PathInfo),
		HasBody:           m.PathInfo.BodyFieldPath != "",
		InputTypeName:     c.methodInOutTypeName(m.InputTypeID, state),
		NameToCamel:       strcase.ToCamel(m.Name),
		NameToPascal:      c.toPascal(m.Name),
		NameToSnake:       strcase.ToSnake(m.Name),
		OutputTypeName:    c.methodInOutTypeName(m.OutputTypeID, state),
		PathParams: mapSlice(PathParams(m, state), func(s *api.Field) *RustField {
			return newRustField(s, c, state)
		}),
		QueryParams: mapSlice(QueryParams(m, state), func(s *api.Field) *RustField {
			return newRustField(s, c, state)
		}),
		IsPageable:          m.IsPageable,
		ServiceNameToPascal: c.toPascal(m.Parent.Name),
		ServiceNameToCamel:  c.toCamel(m.Parent.Name),
		ServiceNameToSnake:  c.toSnake(m.Parent.Name),
		InputTypeID:         m.InputTypeID,
	}
	if m.OperationInfo != nil {
		method.OperationInfo = &RustOperationInfo{
			MetadataType: c.methodInOutTypeName(m.OperationInfo.MetadataTypeID, state),
			ResponseType: c.methodInOutTypeName(m.OperationInfo.ResponseTypeID, state),
		}
	}
	return method
}

func newRustOneOf(oneOf *api.OneOf, c *rustCodec, state *api.APIState) *RustOneOf {
	return &RustOneOf{
		NameToPascal:          c.toPascal(oneOf.Name),
		NameToSnake:           c.toSnake(oneOf.Name),
		NameToSnakeNoMangling: c.toSnakeNoMangling(oneOf.Name),
		FieldType:             c.oneOfType(oneOf, state),
		DocLines:              c.formatDocComments(oneOf.Documentation, state),
		Fields: mapSlice(oneOf.Fields, func(field *api.Field) *RustField {
			return newRustField(field, c, state)
		}),
	}
}

func newRustField(field *api.Field, c *rustCodec, state *api.APIState) *RustField {
	if field == nil {
		return nil
	}
	return &RustField{
		NameToSnake:           c.toSnake(field.Name),
		NameToSnakeNoMangling: c.toSnakeNoMangling(field.Name),
		NameToCamel:           c.toCamel(field.Name),
		NameToPascal:          c.toPascal(field.Name),
		DocLines:              c.formatDocComments(field.Documentation, state),
		FieldAttributes:       c.fieldAttributes(field, state),
		FieldType:             c.fieldType(field, state, false),
		PrimitiveFieldType:    c.fieldType(field, state, true),
		JSONName:              field.JSONName,
		AsQueryParameter:      c.asQueryParameter(field),
	}
}

func newRustEnum(e *api.Enum, c *rustCodec, state *api.APIState) *RustEnum {
	return &RustEnum{
		Name:          c.enumName(e),
		NameSnakeCase: c.toSnake(c.enumName(e)),
		DocLines:      c.formatDocComments(e.Documentation, state),
		Values: mapSlice(e.Values, func(s *api.EnumValue) *RustEnumValue {
			return newRustEnumValue(s, e, c, state)
		}),
	}
}

func newRustEnumValue(ev *api.EnumValue, e *api.Enum, c *rustCodec, state *api.APIState) *RustEnumValue {
	return &RustEnumValue{
		DocLines: c.formatDocComments(ev.Documentation, state),
		Name:     c.enumValueName(ev, state),
		Number:   ev.Number,
		EnumType: c.enumName(e),
	}
}

func rustGeneratedFiles(generateModule bool, state *api.APIState) []GeneratedFile {
	var root string
	switch {
	case generateModule:
		root = "templates/rust/mod"
	case rustHasServices(state):
		root = "templates/rust/nosvc"
	default:
		root = "templates/rust/crate"
	}
	return walkTemplatesDir(rustTemplates, root)
}
