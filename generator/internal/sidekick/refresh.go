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
	"fmt"
	"path"

	"github.com/googleapis/google-cloud-rust/generator/internal/genclient"
	"github.com/googleapis/google-cloud-rust/generator/internal/language"
	"github.com/googleapis/google-cloud-rust/generator/internal/parser"
)

// refresh reruns the generator in one directory, using the configuration
// parameters saved in its `.sidekick.toml` file.
func refresh(rootConfig *Config, cmdLine *CommandLine, output string) error {
	config, err := mergeConfigAndFile(rootConfig, path.Join(output, ".sidekick.toml"))
	if err != nil {
		return err
	}
	if config.General.SpecificationFormat == "" {
		return fmt.Errorf("must provide general.specification-format")
	}
	if config.General.SpecificationSource == "" {
		return fmt.Errorf("must provide general.specification-source")
	}

	specFormat := config.General.SpecificationFormat
	popts := &genclient.ParserOptions{
		Source:        config.General.SpecificationSource,
		ServiceConfig: config.General.ServiceConfig,
		Options:       config.Source,
	}

	var api *genclient.API
	switch specFormat {
	case "openapi":
		api, err = parser.ParseOpenAPI(*popts)
	case "protobuf":
		api, err = parser.ParseProtobuf(*popts)
	default:
		return fmt.Errorf("unknown parser %q", specFormat)
	}
	if err != nil {
		return err
	}

	var (
		codec genclient.LanguageCodec
		copts = &genclient.CodecOptions{
			OutDir:  output,
			Options: config.Codec,
		}
	)
	switch config.General.Language {
	case "rust":
		codec, err = language.NewRustCodec(copts)
	case "go":
		codec, err = language.NewGoCodec(copts)
	default:
		return fmt.Errorf("unknown language: %s", config.General.Language)
	}
	if err != nil {
		return err
	}
	if err := codec.Validate(api); err != nil {
		return err
	}

	request := &genclient.GenerateRequest{
		API:         api,
		Codec:       codec,
		OutDir:      output,
		TemplateDir: config.General.TemplateDir,
	}
	if cmdLine.DryRun {
		return nil
	}
	return genclient.Generate(request)
}
