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

	"github.com/googleapis/google-cloud-rust/generator/internal/api"
)

func newRustCodec(a *api.API, outdir string, options map[string]string) (*rustCodec, error) {
	codec, err := createNewRustCodec(outdir, options)
	if err != nil {
		return nil, err
	}
	if err := rustValidate(codec, a); err != nil {
		return nil, err
	}
	rustLoadWellKnownTypes(a.State)
	return codec, nil
}

func createNewRustCodec(outdir string, options map[string]string) (*rustCodec, error) {
	codec := &rustCodec{
		OutputDirectory:          outdir,
		ModulePath:               "model",
		DeserializeWithdDefaults: true,
		ExtraPackages:            []*rustPackage{},
		PackageMapping:           map[string]*rustPackage{},
	}
	for key, definition := range options {
		switch key {
		case "package-name-override":
			codec.PackageNameOverride = definition
			continue
		case "generate-module":
			value, err := strconv.ParseBool(definition)
			if err != nil {
				return nil, fmt.Errorf("cannot convert `generate-module` value %q to boolean: %w", definition, err)
			}
			codec.GenerateModule = value
			continue
		case "module-path":
			codec.ModulePath = definition
		case "deserialize-with-defaults":
			value, err := strconv.ParseBool(definition)
			if err != nil {
				return nil, fmt.Errorf("cannot convert `deserialize-with-defaults` value %q to boolean: %w", definition, err)
			}
			codec.DeserializeWithdDefaults = value
			continue
		}
		if !strings.HasPrefix(key, "package:") {
			continue
		}
		var specificationPackages []string
		pkg := &rustPackage{
			Name: strings.TrimPrefix(key, "package:"),
		}
		for _, element := range strings.Split(definition, ",") {
			s := strings.SplitN(element, "=", 2)
			if len(s) != 2 {
				return nil, fmt.Errorf("the definition for package %q should be a comma-separated list of key=value pairs, got=%q", key, definition)
			}
			switch s[0] {
			case "package":
				pkg.Package = s[1]
			case "path":
				pkg.Path = s[1]
			case "version":
				pkg.Version = s[1]
			case "source":
				specificationPackages = append(specificationPackages, s[1])
			case "feature":
				pkg.Features = append(pkg.Features, strings.Split(s[1], ",")...)
			case "ignore":
				value, err := strconv.ParseBool(s[1])
				if err != nil {
					return nil, fmt.Errorf("cannot convert `ignore` value %q (part of %q) to boolean: %w", definition, s[1], err)
				}
				pkg.Ignore = value
			default:
				return nil, fmt.Errorf("unknown field %q in definition of rust package %q, got=%q", s[0], key, definition)
			}
		}
		if !pkg.Ignore && pkg.Package == "" {
			return nil, fmt.Errorf("missing rust package name for package %s, got=%s", key, definition)
		}
		codec.ExtraPackages = append(codec.ExtraPackages, pkg)
		for _, source := range specificationPackages {
			codec.PackageMapping[source] = pkg
		}
	}
	return codec, nil
}

type rustCodec struct {
	// The output directory relative to the project root.
	OutputDirectory string
	// Package name override. If not empty, overrides the default package name.
	PackageNameOverride string
	// Generate a module of a larger crate, as opposed to a full crate.
	GenerateModule bool
	// The full path of the generated module within the crate. This defaults to
	// `model`. When generating only a module within a larger crate (see
	// `GenerateModule`), this overrides the path for elements within the crate.
	// Note that using `self` does not work, as the generated code may contain
	// nested modules for nested messages.
	ModulePath string
	// If true, the deserialization functions will accept default values in
	// messages. In almost all cases this should be `true`, but
	DeserializeWithdDefaults bool
	// Additional Rust packages imported by this module. The Mustache template
	// hardcodes a number of packages, but some are configured via the
	// command-line.
	ExtraPackages []*rustPackage
	// A mapping between the specification package names (typically Protobuf),
	// and the Rust package name that contains these types.
	PackageMapping map[string]*rustPackage
	// The source package name (e.g. google.iam.v1 in Protobuf). The codec can
	// generate code for one source package at a time.
	SourceSpecificationPackageName string
}

type rustPackage struct {
	// The name we import this package under.
	Name string
	// If true, ignore the package. We anticipate that the top-level
	// `.sidekick.toml` file will have a number of pre-configured dependencies,
	// but these will be ignored by a handful of packages.
	Ignore bool
	// What the Rust package calls itself.
	Package string
	// The path to file the package locally, unused if empty.
	Path string
	// The version of the package, unused if empty.
	Version string
	// Optional features enabled for the package.
	Features []string
}

var typeMap = map[api.Typez]string{
	api.DOUBLE_TYPE:   "f64",
	api.FLOAT_TYPE:    "f32",
	api.INT64_TYPE:    "i64",
	api.UINT64_TYPE:   "u64",
	api.INT32_TYPE:    "i32",
	api.FIXED64_TYPE:  "u64",
	api.FIXED32_TYPE:  "u32",
	api.BOOL_TYPE:     "bool",
	api.STRING_TYPE:   "String",
	api.BYTES_TYPE:    "bytes::Bytes",
	api.UINT32_TYPE:   "u32",
	api.SFIXED32_TYPE: "i32",
	api.SFIXED64_TYPE: "i64",
	api.SINT32_TYPE:   "i32",
	api.SINT64_TYPE:   "i64",
}

func scalarFieldType(f *api.Field) string {
	out, ok := typeMap[f.Typez]
	if !ok {
		return ""
	}
	return out
}

var rustKeywords = map[string]bool{
	"as":       true,
	"break":    true,
	"const":    true,
	"continue": true,
	"crate":    true,
	"else":     true,
	"enum":     true,
	"extern":   true,
	"false":    true,
	"fn":       true,
	"for":      true,
	"if":       true,
	"impl":     true,
	"in":       true,
	"let":      true,
	"loop":     true,
	"match":    true,
	"mod":      true,
	"move":     true,
	"mut":      true,
	"pub":      true,
	"ref":      true,
	"return":   true,
	"self":     true,
	"Self":     true,
	"static":   true,
	"struct":   true,
	"super":    true,
	"trait":    true,
	"true":     true,
	"type":     true,
	"unsafe":   true,
	"use":      true,
	"where":    true,
	"while":    true,

	// Keywords in Rust 2018+.
	"async": true,
	"await": true,
	"dyn":   true,

	// Reserved
	"abstract": true,
	"become":   true,
	"box":      true,
	"do":       true,
	"final":    true,
	"macro":    true,
	"override": true,
	"priv":     true,
	"typeof":   true,
	"unsized":  true,
	"virtual":  true,
	"yield":    true,

	// Reserved in Rust 2018+
	"try": true,
}

// The list of Rust keywords and reserved words can be found at:
//
//	https://doc.rust-lang.org/reference/keywords.html
func rustEscapeKeyword(symbol string) string {
	_, ok := rustKeywords[symbol]
	if !ok {
		return symbol
	}
	return "r#" + symbol
}
