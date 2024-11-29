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

package api

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/googleapis/google-cloud-rust/generator/internal/testing/sample"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/genproto/googleapis/api/serviceconfig"
	"google.golang.org/protobuf/types/known/apipb"
)

func TestReadServiceConfig(t *testing.T) {
	const serviceConfigPath = "../testing/testdata/googleapis/google/cloud/secretmanager/v1/secretmanager_v1.yaml"
	got, err := readServiceConfig(serviceConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(sample.ServiceConfig, got,
		cmpopts.IgnoreFields(serviceconfig.Service{}, "ConfigVersion", "Publishing", "Authentication"),
		cmpopts.IgnoreUnexported(annotations.HttpRule{}),
		cmpopts.IgnoreUnexported(annotations.Http{}),
		cmpopts.IgnoreUnexported(apipb.Api{}),
		cmpopts.IgnoreUnexported(serviceconfig.AuthenticationRule{}),
		cmpopts.IgnoreUnexported(serviceconfig.Authentication{}),
		cmpopts.IgnoreUnexported(serviceconfig.BackendRule{}),
		cmpopts.IgnoreUnexported(serviceconfig.Backend{}),
		cmpopts.IgnoreUnexported(serviceconfig.DocumentationRule{}),
		cmpopts.IgnoreUnexported(serviceconfig.Documentation{}),
		cmpopts.IgnoreUnexported(serviceconfig.OAuthRequirements{}),
		cmpopts.IgnoreUnexported(serviceconfig.Service{}),
	); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}