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

package sidekick

import (
	"strings"

	"github.com/googleapis/google-cloud-rust/generator/internal/api"
	"github.com/googleapis/google-cloud-rust/generator/internal/language"
	"github.com/iancoleman/strcase"
)

// newTemplateData returns a struct that is used as input to the mustache
// templates. Methods on the types defined in this file are directly associated
// with the mustache tags. For instances the mustache tag {{#Services}} calls
// the [templateData.Services] method. templateData uses the raw input of the
// [API] and uses a [lang.Codec] to transform the input into language
// idiomatic representations.
func newTemplateData(model *api.API, codec language.Codec) *templateData {
	codec.LoadWellKnownTypes(model.State)
	return &templateData{
		s: model,
		c: codec,
	}
}

type templateData struct {
	s *api.API
	c language.Codec
}

func (t *templateData) Name() string {
	return t.s.Name
}

func (t *templateData) Title() string {
	return t.s.Title
}

func (t *templateData) Description() string {
	return t.s.Description
}

func (t *templateData) PackageName() string {
	return t.c.PackageName(t.s)
}

func (t *templateData) RequiredPackages() []string {
	return t.c.RequiredPackages()
}

func (t *templateData) HasServices() bool {
	return len(t.s.Services) > 0
}

func (t *templateData) CopyrightYear() string {
	return t.c.CopyrightYear()
}

func (*templateData) BoilerPlate() []string {
	return append(licenseHeaderBulk(),
		// Mark the code generated from templates as such, and warn reader to
		// not edit the file.
		"",
		" Code generated by sidekick. DO NOT EDIT.")
}

func (t templateData) Imports() []string {
	return t.c.Imports()
}

func (t templateData) DefaultHost() string {
	// APIs, as we generate them today, can only have host. It is true an API
	// can contain many services, but these will all be using the same host.
	// Logic elsewhere asserts this is true.
	if len(t.s.Services) > 0 {
		return t.s.Services[0].DefaultHost
	}
	return ""
}

func (t *templateData) Services() []*service {
	return mapSlice(t.s.Services, func(s *api.Service) *service {
		return &service{
			s:     s,
			c:     t.c,
			state: t.s.State,
		}
	})
}

func (t *templateData) Messages() []*message {
	return mapSlice(t.s.Messages, func(m *api.Message) *message {
		return &message{
			s:     m,
			c:     t.c,
			state: t.s.State,
		}
	})
}

func (t *templateData) NameToLower() string {
	return strings.ToLower(t.s.Name)
}

// service represents a service in an API.
type service struct {
	s     *api.Service
	c     language.Codec
	state *api.APIState
}

func (s *service) Methods() []*method {
	return mapSlice(s.s.Methods, func(m *api.Method) *method {
		return &method{
			s:     m,
			c:     s.c,
			state: s.state,
		}
	})
}

// NameToSnake converts Name to snake_case.
func (s *service) NameToSnake() string {
	return s.c.ToSnake(s.s.Name)
}

// NameToPascanl converts a Name to PascalCase.
func (s *service) NameToPascal() string {
	return s.ServiceNameToPascal()
}

// NameToPascal converts a Name to PascalCase.
func (s *service) ServiceNameToPascal() string {
	return s.c.ToPascal(s.s.Name)
}

// NameToCamel coverts Name to camelCase
func (s *service) NameToCamel() string {
	return s.c.ToCamel(s.s.Name)
}

func (s *service) ServiceName() string {
	return s.s.Name
}

func (s *service) DocLines() []string {
	return s.c.FormatDocComments(s.s.Documentation)
}

func (s *service) DefaultHost() string {
	return s.s.DefaultHost
}

// method defines a RPC belonging to a Service.
type method struct {
	s     *api.Method
	c     language.Codec
	state *api.APIState
}

// NameToSnake converts a Name to snake_case.
func (m *method) NameToSnake() string {
	return strcase.ToSnake(m.s.Name)
}

// NameToCamel converts a Name to camelCase.
func (m *method) NameToCamel() string {
	return strcase.ToCamel(m.s.Name)
}

func (m *method) NameToPascal() string {
	return m.c.ToPascal(m.s.Name)
}

func (m *method) DocLines() []string {
	return m.c.FormatDocComments(m.s.Documentation)
}

func (m *method) InputTypeName() string {
	return m.c.MethodInOutTypeName(m.s.InputTypeID, m.state)
}

func (m *method) OutputTypeName() string {
	return m.c.MethodInOutTypeName(m.s.OutputTypeID, m.state)
}

func (m *method) HTTPMethod() string {
	return m.s.PathInfo.Verb
}

func (m *method) HTTPMethodToLower() string {
	return strings.ToLower(m.s.PathInfo.Verb)
}

func (m *method) HTTPPathFmt() string {
	return m.c.HTTPPathFmt(m.s.PathInfo, m.state)
}

func (m *method) HTTPPathArgs() []string {
	return m.c.HTTPPathArgs(m.s.PathInfo, m.state)
}

func (m *method) QueryParams() []*Field {
	return mapSlice(m.c.QueryParams(m.s, m.state), func(s *api.Field) *Field {
		return newField(s, m.c, m.state)
	})
}

func (m *method) HasBody() bool {
	return m.s.PathInfo.BodyFieldPath != ""
}

func (m *method) BodyAccessor() string {
	return m.c.BodyAccessor(m.s, m.state)
}

// message defines a message used in request or response handling.
type message struct {
	s     *api.Message
	c     language.Codec
	state *api.APIState
}

func (m *message) Fields() []*Field {
	return mapSlice(m.s.Fields, func(s *api.Field) *Field {
		return newField(s, m.c, m.state)
	})
}

// BasicFields returns all fields associated with a message that are not apart
// of a explicit one-ofs.
func (m *message) BasicFields() []*Field {
	filtered := filterSlice(m.s.Fields, func(s *api.Field) bool {
		return !s.IsOneOf
	})
	return mapSlice(filtered, func(s *api.Field) *Field {
		return newField(s, m.c, m.state)
	})
}

// ExplicitOneOfs returns a slice of all explicit one-ofs. Notably this leaves
// out proto3 optional fields which are all considered one-ofs in proto.
func (m *message) ExplicitOneOfs() []*OneOf {
	return mapSlice(m.s.OneOfs, func(s *api.OneOf) *OneOf {
		return newOneOf(s, m.c, m.state)
	})
}

func (m *message) NestedMessages() []*message {
	return mapSlice(m.s.Messages, func(s *api.Message) *message {
		return &message{
			s:     s,
			c:     m.c,
			state: m.state,
		}
	})
}

func (m *message) Enums() []*Enum {
	return mapSlice(m.s.Enums, func(s *api.Enum) *Enum {
		return newEnum(s, m.c, m.state)
	})
}

func (m *message) MessageAttributes() []string {
	return m.c.MessageAttributes(m.s, m.state)
}

func (m *message) Name() string {
	return m.c.MessageName(m.s, m.state)
}

func (m *message) QualifiedName() string {
	return m.c.FQMessageName(m.s, m.state)
}

func (m *message) NameSnakeCase() string {
	return m.c.ToSnake(m.s.Name)
}

// HasNestedTypes returns true if the message has nested types, enums, or
// explicit one-ofs.
func (m *message) HasNestedTypes() bool {
	if len(m.s.Enums) > 0 || len(m.s.OneOfs) > 0 {
		return true
	}
	for _, child := range m.s.Messages {
		if !child.IsMap {
			return true
		}
	}
	return false
}

func (m *message) DocLines() []string {
	return m.c.FormatDocComments(m.s.Documentation)
}

func (m *message) IsMap() bool {
	return m.s.IsMap
}

func filterSlice[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
func mapSlice[T, R any](s []T, f func(T) R) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

type OneOf struct {
	NameToPascal          string
	NameToSnake           string
	NameToSnakeNoMangling string
	FieldType             string
	DocLines              []string
	Fields                []*Field
}

type Field struct {
	NameToSnake           string
	NameToSnakeNoMangling string
	NameToCamel           string
	NameToPascal          string
	DocLines              []string
	FieldAttributes       []string
	FieldType             string
	JSONName              string
	AsQueryParameter      string
}

type Enum struct {
	Name          string
	NameSnakeCase string
	DocLines      []string
	Values        []*EnumValue
}

type EnumValue struct {
	DocLines []string
	Name     string
	Number   int32
	EnumType string
}

func newOneOf(oneOf *api.OneOf, c language.Codec, state *api.APIState) *OneOf {
	return &OneOf{
		NameToPascal:          c.ToPascal(oneOf.Name),
		NameToSnake:           c.ToSnake(oneOf.Name),
		NameToSnakeNoMangling: c.ToSnakeNoMangling(oneOf.Name),
		FieldType:             c.OneOfType(oneOf, state),
		DocLines:              c.FormatDocComments(oneOf.Documentation),
		Fields: mapSlice(oneOf.Fields, func(field *api.Field) *Field {
			return newField(field, c, state)
		}),
	}
}

// Constructor function for c.Field
func newField(field *api.Field, c language.Codec, state *api.APIState) *Field {
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

func newEnum(e *api.Enum, c language.Codec, state *api.APIState) *Enum {
	return &Enum{
		Name:          c.EnumName(e, state),
		NameSnakeCase: c.ToSnake(c.EnumName(e, state)),
		DocLines:      c.FormatDocComments(e.Documentation),
		Values: mapSlice(e.Values, func(s *api.EnumValue) *EnumValue {
			return newEnumValue(s, e, c, state)
		}),
	}
}

// Constructor function for c.EnumValue
func newEnumValue(ev *api.EnumValue, e *api.Enum, c language.Codec, state *api.APIState) *EnumValue {
	return &EnumValue{
		DocLines: c.FormatDocComments(ev.Documentation),
		Name:     c.EnumValueName(ev, state),
		Number:   ev.Number,
		EnumType: c.EnumName(e, state),
	}
}
