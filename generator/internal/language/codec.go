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

import "github.com/googleapis/google-cloud-rust/generator/internal/api"

// Codec defines the behavior required to support language-specific
// customization for template generation. This interface provides methods
// for rendering field types, formatting comments, handling well-known types,
// and more. It is designed to be implemented by specific language codecs.
type Codec interface {
	// TemplateDir returns the directory path containing the templates.
	TemplateDir() string

	// LoadWellKnownTypes loads information into the API state for handling
	// well-known types in the target language (e.g., formatting `timestamppb`
	// or defining wrappers around operations).
	LoadWellKnownTypes(s *api.APIState)

	// FieldAttributes returns a list of attributes (e.g., annotations) for a
	// field, to be included immediately before its definition.
	FieldAttributes(f *api.Field, state *api.APIState) []string

	// FieldType returns the string representation of the type of a message field.
	FieldType(f *api.Field, state *api.APIState) string

	// AsQueryParameter generates the representation of a field when used
	// as a query parameter in an HTTP request.
	AsQueryParameter(f *api.Field, state *api.APIState) string

	// MethodInOutTypeName generates the name for a message type ID when used
	// as an input or output argument in client methods.
	MethodInOutTypeName(id string, state *api.APIState) string

	// MessageAttributes returns a list of attributes (e.g., annotations)
	// for a message, to be included immediately before its definition.
	MessageAttributes(m *api.Message, state *api.APIState) []string

	// MessageName returns the unqualified name of a message as used in type
	// definitions.
	MessageName(m *api.Message, state *api.APIState) string

	// FQMessageName returns the fully-qualified name of a message, as used
	// in references from other package components.
	FQMessageName(m *api.Message, state *api.APIState) string

	// EnumName returns the unqualified name of an enum type.
	EnumName(e *api.Enum, state *api.APIState) string

	// FQEnumName returns the fully-qualified name of an enum type.
	FQEnumName(e *api.Enum, state *api.APIState) string

	// EnumValueName returns the unqualified name of an enum value.
	EnumValueName(e *api.EnumValue, state *api.APIState) string

	// FQEnumValueName returns the fully-qualified name of an enum value.
	FQEnumValueName(e *api.EnumValue, state *api.APIState) string

	// OneOfType generates the string representation of a "one-of" field type.
	OneOfType(o *api.OneOf, state *api.APIState) string

	// BodyAccessor generates the accessor string for retrieving the body
	// of a request (e.g., `.Body()`).
	BodyAccessor(m *api.Method, state *api.APIState) string

	// HTTPPathFmt returns a format string for adding path arguments to a URL.
	// It aligns with the order and values of arguments from HTTPPathArgs.
	HTTPPathFmt(m *api.PathInfo, state *api.APIState) string

	// HTTPPathArgs generates the representation of path arguments, which
	// align with the format string returned by HTTPPathFmt.
	HTTPPathArgs(h *api.PathInfo, state *api.APIState) []string

	// QueryParams returns key-value pairs of query parameter names and
	// their corresponding accessors.
	QueryParams(m *api.Method, state *api.APIState) []*api.Field

	// ToSnake converts a symbol name to `snake_case`, applying language-specific
	// mangling to avoid reserved word clashes.
	ToSnake(string) string

	// ToSnakeNoMangling converts a symbol name to `snake_case` without mangling,
	// intended for use in contexts where templates handle the mangling.
	ToSnakeNoMangling(string) string

	// ToPascal converts a symbol name to `PascalCase`, applying language-specific
	// mangling to avoid reserved word clashes.
	ToPascal(string) string

	// ToCamel converts a symbol name to `camelCase`, applying language-specific
	// mangling to avoid reserved word clashes.
	ToCamel(string) string

	// FormatDocComments reformats documentation comments according to the
	// target language's style guidelines (e.g., resolving references or adding
	// annotations).
	FormatDocComments(string) []string

	// RequiredPackages returns additional lines to be included in a module file.
	RequiredPackages() []string

	// PackageName returns the package name in the target language.
	PackageName(api *api.API) string

	// Validate validates the API for codec-specific restrictions.
	Validate(api *api.API) error

	// AdditionalContext provides language-specific information to the template engine.
	AdditionalContext() any

	// Imports returns a list of imports to be included in the generated code.
	Imports() []string
}
