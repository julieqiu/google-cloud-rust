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
	"path"
	"sort"
	"strings"
	"unicode"

	"github.com/googleapis/google-cloud-rust/generator/internal/api"
	"github.com/iancoleman/strcase"
)

type RustTemplateData struct {
	TemplateDir      string
	Name             string
	Title            string
	Description      string
	PackageName      string
	RequiredPackages []string
	HasServices      bool
	CopyrightYear    string
	BoilerPlate      []string
	Imports          []string
	DefaultHost      string
	Services         []*RustService
	Messages         []*RustMessage
	NameToLower      string
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

type RustMethod struct {
	NameToSnake       string
	NameToCamel       string
	NameToPascal      string
	DocLines          []string
	InputTypeName     string
	OutputTypeName    string
	HTTPMethod        string
	HTTPMethodToLower string
	HTTPPathFmt       string
	HTTPPathArgs      []string
	QueryParams       []*RustField
	HasBody           bool
	BodyAccessor      string
}

type RustMessage struct {
	Fields            []*RustField
	BasicFields       []*RustField
	ExplicitOneOfs    []*RustOneOf
	NestedMessages    []*RustMessage
	Enums             []*RustEnum
	MessageAttributes []string
	Name              string
	QualifiedName     string
	NameSnakeCase     string
	HasNestedTypes    bool
	DocLines          []string
	IsMap             bool
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

type RustField struct {
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

type RustOneOf struct {
	NameToPascal          string
	NameToSnake           string
	NameToSnakeNoMangling string
	FieldType             string
	DocLines              []string
	Fields                []*RustField
}

func rustFormatDocComments(documentation string) []string {
	inBlockQuote := false
	ss := strings.Split(documentation, "\n")
	for i := range ss {
		if strings.HasSuffix(ss[i], "```") {
			if !inBlockQuote {
				ss[i] = ss[i] + "norust"
			}
			inBlockQuote = !inBlockQuote
		}
		ss[i] = fmt.Sprintf("/// %s", ss[i])
		// nit: remove the trailing whitespace, this is unsightly.
		ss[i] = strings.TrimRightFunc(ss[i], unicode.IsSpace)
	}
	return ss
}

func rustEnumValueName(e *api.EnumValue, _ *api.APIState) string {
	// The Protobuf naming convention is to use SCREAMING_SNAKE_CASE, we do not
	// need to change anything for Rust
	return rustEscapeKeyword(e.Name)
}

func rustBodyAccessor(m *api.Method, state *api.APIState) string {
	if m.PathInfo.BodyFieldPath == "*" {
		// no accessor needed, use the whole request
		return ""
	}
	return "." + rustToSnake(m.PathInfo.BodyFieldPath)
}

func rustHTTPPathFmt(m *api.PathInfo, state *api.APIState) string {
	fmt := ""
	for _, segment := range m.PathTemplate {
		if segment.Literal != nil {
			fmt = fmt + "/" + *segment.Literal
		} else if segment.FieldPath != nil {
			fmt = fmt + "/{}"
		} else if segment.Verb != nil {
			fmt = fmt + ":" + *segment.Verb
		}
	}
	return fmt
}

// Returns a Rust expression to access (and if needed validatre) each path parameter.
//
// In most cases the parameter is a simple string in the form `name`. In those
// cases the field *must* be a thing that can be formatted to a string, and
// a simple "req.name" expression will work file.
//
// In some cases the parameter is a sequence of `.` separated fields, in the
// form: `field0.field1 ... .fieldN`. In that case each field from `field0` to
// `fieldN-1` must be optional (they are all messages), and each must be
// validated.
//
// We use the `gax::path_parameter::PathParameter::required()` helper to perform
// this validation. This function recursively creates an expression, the
// recursion starts with
//
// ```rust
// use gax::path_parameter::PathParameter as PP;
// PP::required(&req.field0)?.field1
// ```
//
// And then builds up:
//
// ```rust
// use gax::path_parameter::PathParameter as PP;
// PP::required(PP::required(&req.field0)?.field1)?.field2
// ```
//
// and so on.
func rustUnwrapFieldPath(components []string, requestAccess string) (string, string) {
	if len(components) == 1 {
		return requestAccess + "." + rustToSnake(components[0]), components[0]
	}
	unwrap, name := rustUnwrapFieldPath(components[0:len(components)-1], "&req")
	last := components[len(components)-1]
	return fmt.Sprintf("gax::path_parameter::PathParameter::required(%s, \"%s\").map_err(Error::other)?.%s", unwrap, name, last), ""
}

func derefFieldPath(fieldPath string) string {
	components := strings.Split(fieldPath, ".")
	unwrap, _ := rustUnwrapFieldPath(components, "req")
	return unwrap
}

func rustHTTPPathArgs(h *api.PathInfo, state *api.APIState) []string {
	var args []string
	for _, arg := range h.PathTemplate {
		if arg.FieldPath != nil {
			args = append(args, derefFieldPath(*arg.FieldPath))
		}
	}
	return args
}

func rustQueryParams(m *api.Method, state *api.APIState) []*api.Field {
	msg, ok := state.MessageByID[m.InputTypeID]
	if !ok {
		slog.Error("unable to lookup request type", "id", m.InputTypeID)
		return nil
	}

	var queryParams []*api.Field
	for _, field := range msg.Fields {
		if !m.PathInfo.QueryParameters[field.Name] {
			continue
		}
		queryParams = append(queryParams, field)
	}
	return queryParams
}

// Convert a name to `snake_case`. The Rust naming conventions use this style
// for modules, fields, and functions.
//
// This type of conversion can easily introduce keywords. Consider
//
//	`ToSnake("True") -> "true"`
func rustToSnake(symbol string) string {
	return rustEscapeKeyword(rustToSnakeNoMangling(symbol))
}

func rustToSnakeNoMangling(symbol string) string {
	if strings.ToLower(symbol) == symbol {
		return symbol
	}
	return strcase.ToSnake(symbol)
}

// Convert a name to `PascalCase`.  Strangley, the `strcase` package calls this
// `ToCamel` while usually `camelCase` starts with a lowercase letter. The
// Rust naming convensions use this style for structs, enums and traits.
//
// This type of conversion rarely introduces keywords. The one example is
//
//	`ToPascal("self") -> "Self"`
func rustToPascal(symbol string) string {
	return rustEscapeKeyword(strcase.ToCamel(symbol))
}

func rustToCamel(symbol string) string {
	return rustEscapeKeyword(strcase.ToLowerCamel(symbol))
}

func rustProjectRoot(outputDirectory string) string {
	if outputDirectory == "" {
		return ""
	}
	rel := ".."
	for range strings.Count(outputDirectory, "/") {
		rel = path.Join(rel, "..")
	}
	return rel
}

func rustRequiredPackages(outputDir string, extraPackages []*rustPackage) []string {
	lines := []string{}
	for _, pkg := range extraPackages {
		if pkg.Ignore {
			continue
		}
		components := []string{}
		if pkg.Version != "" {
			components = append(components, fmt.Sprintf("version = %q", pkg.Version))
		}
		if pkg.Path != "" {
			components = append(components, fmt.Sprintf("path = %q", path.Join(rustProjectRoot(outputDir), pkg.Path)))
		}
		if pkg.Package != "" {
			components = append(components, fmt.Sprintf("package = %q", pkg.Package))
		}
		if len(pkg.Features) > 0 {
			feats := strings.Join(mapSlice(pkg.Features, func(s string) string { return fmt.Sprintf("%q", s) }), ", ")
			components = append(components, fmt.Sprintf("features = [%s]", feats))
		}
		lines = append(lines, fmt.Sprintf("%-10s = { %s }", pkg.Name, strings.Join(components, ", ")))
	}
	sort.Strings(lines)
	return lines
}

func rustPackageName(packageNameOverride string, api *api.API) string {
	if packageNameOverride == "" {
		return ""
	}
	name := strings.TrimPrefix(api.PackageName, "google.cloud.")
	name = strings.TrimPrefix(name, "google.")
	name = strings.ReplaceAll(name, ".", "-")
	if name == "" {
		name = api.Name
	}
	return "gcp-sdk-" + name
}

func rustFieldAttributes(f *api.Field, state *api.APIState) []string {
	attributes := rustFieldBaseAttributes(f)
	switch f.Typez {
	case api.DOUBLE_TYPE,
		api.FLOAT_TYPE,
		api.INT32_TYPE,
		api.FIXED32_TYPE,
		api.BOOL_TYPE,
		api.STRING_TYPE,
		api.UINT32_TYPE,
		api.SFIXED32_TYPE,
		api.SINT32_TYPE,
		api.ENUM_TYPE,
		api.GROUP_TYPE:
		return attributes

	case api.INT64_TYPE,
		api.UINT64_TYPE,
		api.FIXED64_TYPE,
		api.SFIXED64_TYPE,
		api.SINT64_TYPE,
		api.BYTES_TYPE:
		formatter := rustFieldFormatter(f.Typez)
		if f.Optional {
			return append(attributes, fmt.Sprintf(`#[serde_as(as = "Option<%s>")]`, formatter))
		}
		if f.Repeated {
			return append(attributes, fmt.Sprintf(`#[serde_as(as = "Vec<%s>")]`, formatter))
		}
		return append(attributes, fmt.Sprintf(`#[serde_as(as = "%s")]`, formatter))

	case api.MESSAGE_TYPE:
		if message, ok := state.MessageByID[f.TypezID]; ok && message.IsMap {
			attributes = append(attributes, `#[serde(skip_serializing_if = "std::collections::HashMap::is_empty")]`)
			var key, value *api.Field
			for _, f := range message.Fields {
				switch f.Name {
				case "key":
					key = f
				case "value":
					value = f
				default:
				}
			}
			if key == nil || value == nil {
				slog.Error("missing key or value in map field")
				return attributes
			}
			keyFormat := rustFieldFormatter(key.Typez)
			valFormat := rustFieldFormatter(value.Typez)
			if keyFormat == "_" && valFormat == "_" {
				return attributes
			}
			return append(attributes, fmt.Sprintf(`#[serde_as(as = "std::collections::HashMap<%s, %s>")]`, keyFormat, valFormat))
		}
		return rustWrapperFieldAttributes(f, attributes)

	default:
		slog.Error("unexpected field type", "field", *f)
		return attributes
	}
}

func rustFieldType(f *api.Field, state *api.APIState) string {
	if f.IsOneOf {
		return rustWrapOneOfField(f, rustBaseFieldType(f, state))
	}
	if f.Repeated {
		return fmt.Sprintf("Vec<%s>", rustBaseFieldType(f, state))
	}
	if f.Optional {
		return fmt.Sprintf("Option<%s>", rustBaseFieldType(f, state))
	}
	return rustBaseFieldType(f, state)
}

func rustWrapOneOfField(f *api.Field, value string) string {
	if f.Typez == api.MESSAGE_TYPE {
		return fmt.Sprintf("(%s)", value)
	}
	return fmt.Sprintf("{ %s: %s }", rustToSnake(f.Name), value)
}

func rustAsQueryParameter(f *api.Field, state *api.APIState) string {
	if f.Typez == api.MESSAGE_TYPE {
		// Query parameters in nested messages are first converted to a
		// `serde_json::Value`` and then recursively merged into the request
		// query. The conversion to `serde_json::Value` is expensive, but very
		// few requests use nested objects as query parameters. Furthermore,
		// the conversion is skipped if the object field is `None`.`
		return fmt.Sprintf("&serde_json::to_value(&req.%s).map_err(Error::serde)?", rustToSnake(f.Name))
	}
	return fmt.Sprintf("&req.%s", rustToSnake(f.Name))
}

func rustTemplateDir(generateModule bool) string {
	if generateModule {
		return "rust/mod"
	}
	return "rust/crate"
}

func rustMessageAttributes(deserializeWithdDefaults bool) []string {
	serde := `#[serde(default, rename_all = "camelCase")]`
	if !deserializeWithdDefaults {
		serde = `#[serde(rename_all = "camelCase")]`
	}
	return []string{
		`#[serde_with::serde_as]`,
		`#[derive(Clone, Debug, Default, PartialEq, serde::Deserialize, serde::Serialize)]`,
		serde,
		`#[non_exhaustive]`,
	}
}

func rustMessageName(m *api.Message, state *api.APIState) string {
	return rustToPascal(m.Name)
}

func rustEnumName(e *api.Enum, state *api.APIState) string {
	return rustToPascal(e.Name)
}

func rustFieldFormatter(typez api.Typez) string {
	switch typez {
	case api.INT64_TYPE,
		api.UINT64_TYPE,
		api.FIXED64_TYPE,
		api.SFIXED64_TYPE,
		api.SINT64_TYPE:
		return "serde_with::DisplayFromStr"
	case api.BYTES_TYPE:
		return "serde_with::base64::Base64"
	default:
		return "_"
	}
}

func rustFieldBaseAttributes(f *api.Field) []string {
	if f.Synthetic {
		return []string{`#[serde(skip)]`}
	}
	if rustToCamel(rustToSnake(f.Name)) != f.JSONName {
		return []string{fmt.Sprintf(`#[serde(rename = "%s")]`, f.JSONName)}
	}
	return []string{}
}

func rustWrapperFieldAttributes(f *api.Field, defaultAttributes []string) []string {
	var formatter string
	switch f.TypezID {
	case ".google.protobuf.BytesValue":
		formatter = rustFieldFormatter(api.BYTES_TYPE)
	case ".google.protobuf.UInt64Value":
		formatter = rustFieldFormatter(api.UINT64_TYPE)
	case ".google.protobuf.Int64Value":
		formatter = rustFieldFormatter(api.INT64_TYPE)
	default:
		return defaultAttributes
	}
	return []string{fmt.Sprintf(`#[serde_as(as = "Option<%s>")]`, formatter)}
}

func rustLoadWellKnownTypes(s *api.APIState) {
	// TODO(#77) - replace these placeholders with real types
	wellKnown := []*api.Message{
		{
			ID:      ".google.protobuf.Any",
			Name:    "Any",
			Package: "google.protobuf",
		},
		{
			ID:      ".google.protobuf.Empty",
			Name:    "Empty",
			Package: "google.protobuf",
		},
		{
			ID:      ".google.protobuf.FieldMask",
			Name:    "FieldMask",
			Package: "google.protobuf",
		},
		{
			ID:      ".google.protobuf.Duration",
			Name:    "Duration",
			Package: "google.protobuf",
		},
		{
			ID:      ".google.protobuf.Timestamp",
			Name:    "Timestamp",
			Package: "google.protobuf",
		},
	}
	for _, message := range wellKnown {
		s.MessageByID[message.ID] = message
	}
}

// Constructor function for rustEnum
func newRustEnum(e *api.Enum, c *rustCodec, state *api.APIState) *RustEnum {
	return &RustEnum{
		Name:          rustEnumName(e, state),
		NameSnakeCase: rustToSnake(rustEnumName(e, state)),
		DocLines:      rustFormatDocComments(e.Documentation),
		Values: mapSlice(e.Values, func(s *api.EnumValue) *RustEnumValue {
			return newRustEnumValue(s, e, c, state)
		}),
	}
}

// Constructor function for rustEnumValue
func newRustEnumValue(ev *api.EnumValue, e *api.Enum, c *rustCodec, state *api.APIState) *RustEnumValue {
	return &RustEnumValue{
		DocLines: rustFormatDocComments(ev.Documentation),
		Name:     rustEnumValueName(ev, state),
		Number:   ev.Number,
		EnumType: rustEnumName(e, state),
	}
}

// Constructor function for rustField
func newRustField(field *api.Field, c *rustCodec, state *api.APIState) *RustField {
	return &RustField{
		NameToSnake:           rustToSnake(field.Name),
		NameToSnakeNoMangling: rustToSnakeNoMangling(field.Name),
		NameToCamel:           rustToCamel(field.Name),
		NameToPascal:          rustToPascal(field.Name),
		DocLines:              rustFormatDocComments(field.Documentation),
		FieldAttributes:       rustFieldAttributes(field, state),
		FieldType:             rustFieldType(field, state),
		JSONName:              field.JSONName,
		AsQueryParameter:      rustAsQueryParameter(field, state),
	}
}

// Constructor function for rustOneOf
func newRustOneOf(oneOf *api.OneOf, c *rustCodec, state *api.APIState) *RustOneOf {
	return &RustOneOf{
		NameToPascal:          rustToPascal(oneOf.Name),
		NameToSnake:           rustToSnake(oneOf.Name),
		NameToSnakeNoMangling: rustToSnakeNoMangling(oneOf.Name),
		FieldType:             rustOneOfType(c, oneOf, state),
		DocLines:              rustFormatDocComments(oneOf.Documentation),
		Fields: mapSlice(oneOf.Fields, func(field *api.Field) *RustField {
			return newRustField(field, c, state)
		}),
	}
}

func newRustMessage(m *api.Message, c *rustCodec, state *api.APIState) *RustMessage {
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
		MessageAttributes: rustMessageAttributes(c.DeserializeWithdDefaults),
		Name:              rustMessageName(m, state),
		QualifiedName:     rustFQMessageName(c, m),
		NameSnakeCase:     rustToSnake(m.Name),
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
		DocLines: rustFormatDocComments(m.Documentation),
		IsMap:    m.IsMap,
	}
}

// Constructor function for rustService
func newRustService(s *api.Service, c *rustCodec, state *api.APIState) *RustService {
	return &RustService{
		Methods: mapSlice(s.Methods, func(m *api.Method) *RustMethod {
			return newRustMethod(m, c, state)
		}),
		NameToSnake:         rustToSnake(s.Name),
		NameToPascal:        rustToPascal(s.Name),
		ServiceNameToPascal: rustToPascal(s.Name), // Alias for clarity
		NameToCamel:         rustToCamel(s.Name),
		ServiceName:         s.Name,
		DocLines:            rustFormatDocComments(s.Documentation),
		DefaultHost:         s.DefaultHost,
	}
}
