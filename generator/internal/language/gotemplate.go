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
	"strings"

	"github.com/googleapis/google-cloud-rust/generator/internal/api"
	"github.com/googleapis/google-cloud-rust/generator/internal/license"
	"github.com/iancoleman/strcase"
)

type GoTemplateData struct {
	Name              string
	Title             string
	Description       string
	PackageName       string
	SourcePackageName string
	PackageVersion    string
	RequiredPackages  []string
	HasServices       bool
	CopyrightYear     string
	BoilerPlate       []string
	Imports           []string
	DefaultHost       string
	Services          []*GoService
	Messages          []*GoMessage
	Enums             []*GoEnum
	NameToLower       string
	NotForPublication bool
	GoPackage         string
}

type GoService struct {
	Methods             []*GoMethod
	NameToSnake         string
	NameToPascal        string
	ServiceNameToPascal string
	NameToCamel         string
	ServiceName         string
	DocLines            []string
	DefaultHost         string
}

type GoMessage struct {
	Fields             []*GoField
	BasicFields        []*GoField
	ExplicitOneOfs     []*GoOneOf
	NestedMessages     []*GoMessage
	Enums              []*GoEnum
	Name               string
	QualifiedName      string
	NameSnakeCase      string
	HasNestedTypes     bool
	DocLines           []string
	IsMap              bool
	IsPageableResponse bool
	PageableItem       *GoField
	ID                 string
	// The FQN is the source specification
	SourceFQN string
	// If true, this is a synthetic message, some generation is skipped for
	// synthetic messages
	HasSyntheticFields bool
}

type GoMethod struct {
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
	PathParams          []*GoField
	QueryParams         []*GoField
	HasBody             bool
	BodyAccessor        string
	IsPageable          bool
	ServiceNameToPascal string
	ServiceNameToCamel  string
	ServiceNameToSnake  string
	InputTypeID         string
	InputType           *GoMessage
	OperationInfo       *GoOperationInfo
}

type GoOperationInfo struct {
	MetadataType string
	ResponseType string
}

type GoOneOf struct {
	NameToPascal          string
	NameToSnake           string
	NameToSnakeNoMangling string
	DocLines              []string
	Fields                []*GoField
}

type GoField struct {
	NameToSnake           string
	NameToSnakeNoMangling string
	NameToCamel           string
	NameToPascal          string
	DocLines              []string
	FieldType             string
	PrimitiveFieldType    string
	JSONName              string
	AsQueryParameter      string
}

type GoEnum struct {
	Name          string
	NameSnakeCase string
	DocLines      []string
	Values        []*GoEnumValue
}

type GoEnumValue struct {
	DocLines []string
	Name     string
	Number   int32
	EnumType string
}

// newGoTemplateData creates a struct used as input for Mustache templates.
// Fields and methods defined in this struct directly correspond to Mustache
// tags. For example, the Mustache tag {{#Services}} uses the
// [Template.Services] field.
func newGoTemplateData(model *api.API, c *goCodec) *GoTemplateData {
	c.loadWellKnownTypes(model.State)
	data := &GoTemplateData{
		Name:              model.Name,
		Title:             model.Title,
		Description:       model.Description,
		PackageName:       c.packageName(model),
		SourcePackageName: c.sourcePackageName(),
		PackageVersion:    c.packageVersion(),
		HasServices:       len(model.Services) > 0,
		CopyrightYear:     c.copyrightYear(),
		BoilerPlate: append(license.LicenseHeaderBulk(),
			"",
			" Code generated by sidekick. DO NOT EDIT."),
		Imports: c.imports(),
		DefaultHost: func() string {
			if len(model.Services) > 0 {
				return model.Services[0].DefaultHost
			}
			return ""
		}(),
		Services: mapSlice(model.Services, func(s *api.Service) *GoService {
			return newGoService(s, c, model.State)
		}),
		Messages: mapSlice(model.Messages, func(m *api.Message) *GoMessage {
			return newGoMessage(m, c, model.State)
		}),
		Enums: mapSlice(model.Enums, func(e *api.Enum) *GoEnum {
			return newGoEnum(e, c, model.State)
		}),
		NameToLower:       strings.ToLower(model.Name),
		NotForPublication: c.doNotPublish,
		GoPackage:         c.goPackageName,
	}
	// Delay this until the *GoCodec had a chance to compute what packages are
	// used.
	data.RequiredPackages = c.requiredPackages()

	messagesByID := map[string]*GoMessage{}
	for _, m := range data.Messages {
		messagesByID[m.ID] = m
	}
	for _, s := range data.Services {
		for _, method := range s.Methods {
			if msg, ok := messagesByID[method.InputTypeID]; ok {
				method.InputType = msg
			} else if m, ok := model.State.MessageByID[method.InputTypeID]; ok {
				method.InputType = newGoMessage(m, c, model.State)
			}
		}
	}
	return data
}

func newGoService(s *api.Service, c *goCodec, state *api.APIState) *GoService {
	// Some codecs skip some methods.
	methods := filterSlice(s.Methods, func(m *api.Method) bool {
		return c.generateMethod(m)
	})
	return &GoService{
		Methods: mapSlice(methods, func(m *api.Method) *GoMethod {
			return newGoMethod(m, c, state)
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

func newGoMessage(m *api.Message, c *goCodec, state *api.APIState) *GoMessage {
	hasSyntheticFields := false
	for _, f := range m.Fields {
		if f.Synthetic {
			hasSyntheticFields = true
			break
		}
	}
	return &GoMessage{
		Fields: mapSlice(m.Fields, func(s *api.Field) *GoField {
			return newGoField(s, c, state)
		}),
		BasicFields: func() []*GoField {
			filtered := filterSlice(m.Fields, func(s *api.Field) bool {
				return !s.IsOneOf
			})
			return mapSlice(filtered, func(s *api.Field) *GoField {
				return newGoField(s, c, state)
			})
		}(),
		ExplicitOneOfs: mapSlice(m.OneOfs, func(s *api.OneOf) *GoOneOf {
			return newGoOneOf(s, c, state)
		}),
		NestedMessages: mapSlice(m.Messages, func(s *api.Message) *GoMessage {
			return newGoMessage(s, c, state)
		}),
		Enums: mapSlice(m.Enums, func(s *api.Enum) *GoEnum {
			return newGoEnum(s, c, state)
		}),
		Name:          c.messageName(m),
		QualifiedName: c.messageName(m),
		NameSnakeCase: c.toSnake(m.Name),
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
		PageableItem:       newGoField(m.PageableItem, c, state),
		ID:                 m.ID,
		SourceFQN:          strings.TrimPrefix(m.ID, "."),
		HasSyntheticFields: hasSyntheticFields,
	}
}

func newGoMethod(m *api.Method, c *goCodec, state *api.APIState) *GoMethod {
	method := &GoMethod{
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
		PathParams: mapSlice(PathParams(m, state), func(s *api.Field) *GoField {
			return newGoField(s, c, state)
		}),
		QueryParams: mapSlice(QueryParams(m, state), func(s *api.Field) *GoField {
			return newGoField(s, c, state)
		}),
		IsPageable:          m.IsPageable,
		ServiceNameToPascal: c.toPascal(m.Parent.Name),
		ServiceNameToCamel:  c.toCamel(m.Parent.Name),
		ServiceNameToSnake:  c.toSnake(m.Parent.Name),
		InputTypeID:         m.InputTypeID,
	}
	if m.OperationInfo != nil {
		method.OperationInfo = &GoOperationInfo{
			MetadataType: c.methodInOutTypeName(m.OperationInfo.MetadataTypeID, state),
			ResponseType: c.methodInOutTypeName(m.OperationInfo.ResponseTypeID, state),
		}
	}
	return method
}

func newGoOneOf(oneOf *api.OneOf, c *goCodec, state *api.APIState) *GoOneOf {
	return &GoOneOf{
		NameToPascal:          c.toPascal(oneOf.Name),
		NameToSnake:           c.toSnake(oneOf.Name),
		NameToSnakeNoMangling: c.toSnakeNoMangling(oneOf.Name),
		DocLines:              c.formatDocComments(oneOf.Documentation, state),
		Fields: mapSlice(oneOf.Fields, func(field *api.Field) *GoField {
			return newGoField(field, c, state)
		}),
	}
}

func newGoField(field *api.Field, c *goCodec, state *api.APIState) *GoField {
	if field == nil {
		return nil
	}
	return &GoField{
		NameToSnake:           c.toSnake(field.Name),
		NameToSnakeNoMangling: c.toSnakeNoMangling(field.Name),
		NameToCamel:           c.toCamel(field.Name),
		NameToPascal:          c.toPascal(field.Name),
		DocLines:              c.formatDocComments(field.Documentation, state),
		FieldType:             c.fieldType(field, state),
		PrimitiveFieldType:    c.primitiveFieldType(field, state),
		JSONName:              field.JSONName,
		AsQueryParameter:      c.asQueryParameter(field, state),
	}
}

func newGoEnum(e *api.Enum, c *goCodec, state *api.APIState) *GoEnum {
	return &GoEnum{
		Name:          c.enumName(e),
		NameSnakeCase: c.toSnake(c.enumName(e)),
		DocLines:      c.formatDocComments(e.Documentation, state),
		Values: mapSlice(e.Values, func(s *api.EnumValue) *GoEnumValue {
			return newGoEnumValue(s, e, c, state)
		}),
	}
}

func newGoEnumValue(ev *api.EnumValue, e *api.Enum, c *goCodec, state *api.APIState) *GoEnumValue {
	return &GoEnumValue{
		DocLines: c.formatDocComments(ev.Documentation, state),
		Name:     c.enumValueName(ev),
		Number:   ev.Number,
		EnumType: c.enumName(e),
	}
}