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

// Package genclient is a Schema and Language agnostic code generator that applies
// an API model to a mustache template.
package genclient

import (
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/cbroglie/mustache"
)

// LanguageCodec is an adapter used to transform values into language idiomatic
// representations. This is used to manipulate the data the is fed into
// templates to generate clients.
type LanguageCodec interface {
	// TemplateDir returns the directory containing the templates.
	TemplateDir() string
	// LoadWellKnownTypes allows a language to load information into the state
	// for any wellknown types. For example defining how timestamppb should be
	// represented in a given language or wrappers around operations.
	LoadWellKnownTypes(s *APIState)
	// FieldAttributes returns a (possibly empty) list of "attributes" included
	// immediately before the field definition.
	FieldAttributes(f *Field, state *APIState) []string
	// FieldType returns a string representation of a message field type.
	FieldType(f *Field, state *APIState) string
	// The field when used to build the query.
	AsQueryParameter(f *Field, state *APIState) string
	// The name of a message type ID when used as an input or output argument
	// in the client methods.
	MethodInOutTypeName(id string, state *APIState) string
	// Returns a (possibly empty) list of "attributes" included immediately
	// before the message definition.
	MessageAttributes(m *Message, state *APIState) []string
	// The (unqualified) message name, as used when defining the type to
	// represent it.
	MessageName(m *Message, state *APIState) string
	// The fully-qualified message name, as used when referring to the name from
	// another place in the package.
	FQMessageName(m *Message, state *APIState) string
	// The (unqualified) enum name, as used when defining the type to
	// represent it.
	EnumName(e *Enum, state *APIState) string
	// The fully-qualified enum name, as used when referring to the name from
	// another place in the package.
	FQEnumName(e *Enum, state *APIState) string
	// The (unqualified) enum value name, as used when defining the constant,
	// variable, or enum value that holds it.
	EnumValueName(e *EnumValue, state *APIState) string
	// The fully qualified enum value name, as used when using the constant,
	// variable, or enum value that holds it.
	FQEnumValueName(e *EnumValue, state *APIState) string
	// OneOfType returns a string representation of a one-of field type.
	OneOfType(o *OneOf, state *APIState) string
	// BodyAccessor returns a string representation of the accessor used to
	// get the body out of a request. For instance this might return `.Body()`.
	BodyAccessor(m *Method, state *APIState) string
	// HTTPPathFmt returns a format string used for adding path arguments to a
	// URL. The replacements should align in both order and value from what is
	// returned from HTTPPathArgs.
	HTTPPathFmt(m *PathInfo, state *APIState) string
	// HTTPPathArgs returns a string representation of the path arguments. This
	// should be used in conjunction with HTTPPathFmt. An example return value
	// might be `, req.PathParam()`
	HTTPPathArgs(h *PathInfo, state *APIState) []string
	// QueryParams returns key-value pairs of name to accessor for query params.
	// An example return value might be
	// `&Pair{Key: "secretId", Value: "req.SecretId()"}`
	QueryParams(m *Method, state *APIState) []*Field
	// ToSnake converts a symbol name to `snake_case`, applying any mangling
	// required by the language, e.g., to avoid clashes with reserved words.
	ToSnake(string) string
	// ToSnakeNoMangling converts a symbol name to `snake_case`, without any
	// mangling to avoid reserved words. This is useful when the template is
	// already going to mangle the name, e.g., by adding a prefix or suffix.
	// Since the templates are language specific, their authors can determine
	// when to use `ToSnake` or `ToSnakeNoMangling`.
	ToSnakeNoMangling(string) string
	// ToPascal converts a symbol name to `PascalCase`, applying any mangling
	// required by the language, e.g., to avoid clashes with reserved words.
	ToPascal(string) string
	// ToCamel converts a symbol name to `camelCase` (sometimes called
	// "lowercase CamelCase"), applying any mangling required by the language,
	// e.g., to avoid clashes with reserved words.
	ToCamel(string) string
	// Reformat ${Lang}Doc comments according to the language-specific rules.
	// For example,
	//   - The protos in googleapis include cross-references in the format
	//     `[Foo][proto.package.name.Foo]`, this should become links to the
	//     language entities, in the language documentation.
	//   - Rust requires a `norust` annotation in all blockquotes, that is,
	//     any ```-sections. Without this annotation Rustdoc assumes the
	//     blockquote is an Rust code snippet and attempts to compile it.
	FormatDocComments(string) []string
	// Returns a extra set of lines to insert in the module file.
	// The format of these lines is specific to each language.
	RequiredPackages() []string
	// The package name in the destination language. May be empty, some
	// languages do not have a package manager.
	PackageName(api *API) string
	// Validate an API, some codecs impose restrictions on the input API.
	Validate(api *API) error
	// The year when this package was first generated.
	CopyrightYear() string
	// Pass language-specific information from the Codec to the template engine.
	// Prefer using specific methods when the information is applicable to most
	// (or many) languages. Use this method when the information is application
	// to only one language.
	AdditionalContext() any
	// Imports to add.
	Imports() []string
}

type ParserOptions struct {
	// The location where the specification can be found.
	Source string
	// The location of the service configuration file.
	ServiceConfig string
	// Additional options.
	Options map[string]string
}

type CodecOptions struct {
	// The output location within ProjectRoot.
	OutDir string
	// Additional options.
	Options map[string]string
}

// GenerateRequest used to generate clients.
type GenerateRequest struct {
	// The in memory representation of a parsed input.
	API *API
	// An adapter to transform values into language idiomatic representations.
	Codec LanguageCodec
	// OutDir is the path to the output directory.
	OutDir string
	// Template directory
	TemplateDir string
}

func (r *GenerateRequest) outDir() string {
	if r.OutDir == "" {
		wd, _ := os.Getwd()
		return wd
	}
	return r.OutDir
}

// Generate takes some state and applies it to a template to create a client
// library.
func Generate(req *GenerateRequest) error {
	data := newTemplateData(req.API, req.Codec)
	root := filepath.Join(req.TemplateDir, req.Codec.TemplateDir())
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			dn := filepath.Join(req.outDir(), strings.TrimPrefix(path, root))
			os.MkdirAll(dn, 0777) // Ignore errors
			return nil
		}
		if filepath.Ext(path) != ".mustache" {
			return nil
		}
		if strings.Count(d.Name(), ".") == 1 {
			// skipping partials
			return nil
		}
		var context []any
		context = append(context, data)
		if req.Codec.AdditionalContext() != nil {
			context = append(context, req.Codec.AdditionalContext())
		}
		s, err := mustache.RenderFile(path, context...)
		if err != nil {
			return err
		}
		fn := filepath.Join(req.outDir(), filepath.Dir(strings.TrimPrefix(path, root)), strings.TrimSuffix(d.Name(), ".mustache"))
		return os.WriteFile(fn, []byte(s), os.ModePerm)
	})
	if err != nil {
		slog.Error("error walking templates", "err", err.Error())
		return err
	}

	return nil
}
