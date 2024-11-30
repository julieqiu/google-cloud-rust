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
	"github.com/iancoleman/strcase"
)

type GoTemplateData struct {
	GoPackage string
	TemplateData
}

// NewGoTemplateData creates a struct used as input for Mustache templates.
// Fields and methods defined in this struct directly correspond to Mustache
// tags. For example, the Mustache tag {{#Services}} uses the
// [Template.Services] field.
func NewGoTemplateData(model *api.API, c *GoCodec) *GoTemplateData {
	c.LoadWellKnownTypes(model.State)
	return &GoTemplateData{
		GoPackage: c.GoPackageName,
		TemplateData: TemplateData{
			TemplateDir:      c.TemplateDir(),
			Name:             model.Name,
			Title:            model.Title,
			Description:      model.Description,
			PackageName:      c.PackageName(model),
			RequiredPackages: c.RequiredPackages(),
			HasServices:      len(model.Services) > 0,
			Imports:          c.Imports(),
			DefaultHost: func() string {
				if len(model.Services) > 0 {
					return model.Services[0].DefaultHost
				}
				return ""
			}(),
			Services: mapSlice(model.Services, func(s *api.Service) *Service {
				return newGoService(s, c, model.State)
			}),
			Messages: mapSlice(model.Messages, func(m *api.Message) *Message {
				return newGoMessage(m, c, model.State)
			}),
			NameToLower: strings.ToLower(model.Name),
		},
	}
}

func newGoService(s *api.Service, c *GoCodec, state *api.APIState) *Service {
	return &Service{
		Methods: mapSlice(s.Methods, func(m *api.Method) *Method {
			return newGoMethod(m, c, state)
		}),
		NameToSnake:         c.ToSnake(s.Name),
		NameToPascal:        c.ToPascal(s.Name),
		ServiceNameToPascal: c.ToPascal(s.Name), // Alias for clarity
		NameToCamel:         c.ToCamel(s.Name),
		ServiceName:         s.Name,
		DocLines:            c.FormatDocComments(s.Documentation),
		DefaultHost:         s.DefaultHost,
	}
}

func newGoMessage(m *api.Message, c *GoCodec, state *api.APIState) *Message {
	return &Message{
		Fields: mapSlice(m.Fields, func(s *api.Field) *Field {
			return newGoField(s, c, state)
		}),
		BasicFields: func() []*Field {
			filtered := filterSlice(m.Fields, func(s *api.Field) bool {
				return !s.IsOneOf
			})
			return mapSlice(filtered, func(s *api.Field) *Field {
				return newGoField(s, c, state)
			})
		}(),
		ExplicitOneOfs: mapSlice(m.OneOfs, func(s *api.OneOf) *OneOf {
			return newGoOneOf(s, c, state)
		}),
		NestedMessages: mapSlice(m.Messages, func(s *api.Message) *Message {
			return newGoMessage(s, c, state)
		}),
		Enums: mapSlice(m.Enums, func(s *api.Enum) *Enum {
			return newGoEnum(s, c, state)
		}),
		MessageAttributes: c.MessageAttributes(m, state),
		Name:              c.MessageName(m, state),
		QualifiedName:     c.FQMessageName(m, state),
		NameSnakeCase:     c.ToSnake(m.Name),
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
		DocLines: c.FormatDocComments(m.Documentation),
		IsMap:    m.IsMap,
	}
}

func newGoMethod(m *api.Method, c *GoCodec, state *api.APIState) *Method {
	return &Method{
		BodyAccessor:      c.BodyAccessor(m, state),
		DocLines:          c.FormatDocComments(m.Documentation),
		HTTPMethod:        m.PathInfo.Verb,
		HTTPMethodToLower: strings.ToLower(m.PathInfo.Verb),
		HTTPPathArgs:      c.HTTPPathArgs(m.PathInfo, state),
		HTTPPathFmt:       c.HTTPPathFmt(m.PathInfo, state),
		HasBody:           m.PathInfo.BodyFieldPath != "",
		InputTypeName:     c.MethodInOutTypeName(m.InputTypeID, state),
		NameToCamel:       strcase.ToCamel(m.Name),
		NameToPascal:      c.ToPascal(m.Name),
		NameToSnake:       strcase.ToSnake(m.Name),
		OutputTypeName:    c.MethodInOutTypeName(m.OutputTypeID, state),
		QueryParams: mapSlice(c.QueryParams(m, state), func(s *api.Field) *Field {
			return newGoField(s, c, state)
		}),
	}
}

func newGoOneOf(oneOf *api.OneOf, c *GoCodec, state *api.APIState) *OneOf {
	return &OneOf{
		NameToPascal:          c.ToPascal(oneOf.Name),
		NameToSnake:           c.ToSnake(oneOf.Name),
		NameToSnakeNoMangling: c.ToSnakeNoMangling(oneOf.Name),
		FieldType:             c.OneOfType(oneOf, state),
		DocLines:              c.FormatDocComments(oneOf.Documentation),
		Fields: mapSlice(oneOf.Fields, func(field *api.Field) *Field {
			return newGoField(field, c, state)
		}),
	}
}

// Constructor function for c.Field
func newGoField(field *api.Field, c *GoCodec, state *api.APIState) *Field {
	return &Field{
		NameToSnake:           c.ToSnake(field.Name),
		NameToSnakeNoMangling: c.ToSnakeNoMangling(field.Name),
		NameToCamel:           c.ToCamel(field.Name),
		NameToPascal:          c.ToPascal(field.Name),
		DocLines:              c.FormatDocComments(field.Documentation),
		FieldAttributes:       c.FieldAttributes(field, state),
		FieldType:             c.FieldType(field, state),
		JSONName:              field.JSONName,
		AsQueryParameter:      c.AsQueryParameter(field, state),
	}
}

func newGoEnum(e *api.Enum, c *GoCodec, state *api.APIState) *Enum {
	return &Enum{
		Name:          c.EnumName(e, state),
		NameSnakeCase: c.ToSnake(c.EnumName(e, state)),
		DocLines:      c.FormatDocComments(e.Documentation),
		Values: mapSlice(e.Values, func(s *api.EnumValue) *EnumValue {
			return newGoEnumValue(s, e, c, state)
		}),
	}
}

// Constructor function for c.EnumValue
func newGoEnumValue(ev *api.EnumValue, e *api.Enum, c *GoCodec, state *api.APIState) *EnumValue {
	return &EnumValue{
		DocLines: c.FormatDocComments(ev.Documentation),
		Name:     c.EnumValueName(ev, state),
		Number:   ev.Number,
		EnumType: c.EnumName(e, state),
	}
}
