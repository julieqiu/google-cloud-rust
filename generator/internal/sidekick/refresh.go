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

	"github.com/googleapis/google-cloud-rust/generator/internal/api"
	"github.com/googleapis/google-cloud-rust/generator/internal/language"
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

	var a *api.API
	switch config.General.SpecificationFormat {
	case "openapi":
		a, err = api.FromOpenAPI(config.General.SpecificationSource, config.General.ServiceConfig, config.Source)
	case "protobuf":
		a, err = api.FromProtobuf(config.General.SpecificationSource, config.General.ServiceConfig, config.Source)
	default:
		return fmt.Errorf("unknown parser %q", config.General.SpecificationFormat)
	}
	if err != nil {
		return err
	}

	root, context, err := language.NewTemplateData(a, config.Codec, output, config.General.Language, config.General.TemplateDir)
	if err != nil {
		return err
	}

	request := &generateClientRequest{
		API:    a,
		OutDir: output,
	}
	if cmdLine.DryRun {
		return nil
	}
	return generateClient(root, request, context)
}
