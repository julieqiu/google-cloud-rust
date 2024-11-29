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
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/genproto/googleapis/api/serviceconfig"
	"google.golang.org/protobuf/types/known/apipb"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestProtobuf_Info(t *testing.T) {
	var serviceConfig = &serviceconfig.Service{
		Name:  "secretmanager.googleapis.com",
		Title: "Secret Manager API",
		Documentation: &serviceconfig.Documentation{
			Summary:  "Stores sensitive data such as API keys, passwords, and certificates.\nProvides convenience while improving security.",
			Overview: "Secret Manager Overview",
		},
	}

	test := makeAPIForProtobuf(serviceConfig, newTestCodeGeneratorRequest(t, "scalar.proto"))
	if test.Name != "secretmanager" {
		t.Errorf("want = %q; got = %q", "secretmanager", test.Name)
	}
	if test.Title != serviceConfig.Title {
		t.Errorf("want = %q; got = %q", serviceConfig.Title, test.Name)
	}
	if diff := cmp.Diff(test.Description, serviceConfig.Documentation.Summary); diff != "" {
		t.Errorf("description mismatch (-want, +got):\n%s", diff)
	}
}

func TestProtobuf_Scalar(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "scalar.proto"))
	message, ok := test.State.MessageByID[".test.Fake"]
	if !ok {
		t.Fatalf("Cannot find message %s in API State", ".test.Fake")
	}
	checkProtobufMessage(t, *message, Message{
		Name:          "Fake",
		Package:       "test",
		ID:            ".test.Fake",
		Documentation: "A test message.",
		Fields: []*Field{
			{
				Documentation: "A singular field tag = 1",
				Name:          "f_double",
				JSONName:      "fDouble",
				ID:            ".test.Fake.f_double",
				Typez:         DOUBLE_TYPE,
			},
			{
				Documentation: "A singular field tag = 2",
				Name:          "f_float",
				JSONName:      "fFloat",
				ID:            ".test.Fake.f_float",
				Typez:         FLOAT_TYPE,
			},
			{
				Documentation: "A singular field tag = 3",
				Name:          "f_int64",
				JSONName:      "fInt64",
				ID:            ".test.Fake.f_int64",
				Typez:         INT64_TYPE,
			},
			{
				Documentation: "A singular field tag = 4",
				Name:          "f_uint64",
				JSONName:      "fUint64",
				ID:            ".test.Fake.f_uint64",
				Typez:         UINT64_TYPE,
			},
			{
				Documentation: "A singular field tag = 5",
				Name:          "f_int32",
				JSONName:      "fInt32",
				ID:            ".test.Fake.f_int32",
				Typez:         INT32_TYPE,
			},
			{
				Documentation: "A singular field tag = 6",
				Name:          "f_fixed64",
				JSONName:      "fFixed64",
				ID:            ".test.Fake.f_fixed64",
				Typez:         FIXED64_TYPE,
			},
			{
				Documentation: "A singular field tag = 7",
				Name:          "f_fixed32",
				JSONName:      "fFixed32",
				ID:            ".test.Fake.f_fixed32",
				Typez:         FIXED32_TYPE,
			},
			{
				Documentation: "A singular field tag = 8",
				Name:          "f_bool",
				JSONName:      "fBool",
				ID:            ".test.Fake.f_bool",
				Typez:         BOOL_TYPE,
			},
			{
				Documentation: "A singular field tag = 9",
				Name:          "f_string",
				JSONName:      "fString",
				ID:            ".test.Fake.f_string",
				Typez:         STRING_TYPE,
			},
			{
				Documentation: "A singular field tag = 12",
				Name:          "f_bytes",
				JSONName:      "fBytes",
				ID:            ".test.Fake.f_bytes",
				Typez:         BYTES_TYPE,
			},
			{
				Documentation: "A singular field tag = 13",
				Name:          "f_uint32",
				JSONName:      "fUint32",
				ID:            ".test.Fake.f_uint32",
				Typez:         UINT32_TYPE,
			},
			{
				Documentation: "A singular field tag = 15",
				Name:          "f_sfixed32",
				JSONName:      "fSfixed32",
				ID:            ".test.Fake.f_sfixed32",
				Typez:         SFIXED32_TYPE,
			},
			{
				Documentation: "A singular field tag = 16",
				Name:          "f_sfixed64",
				JSONName:      "fSfixed64",
				ID:            ".test.Fake.f_sfixed64",
				Typez:         SFIXED64_TYPE,
			},
			{
				Documentation: "A singular field tag = 17",
				Name:          "f_sint32",
				JSONName:      "fSint32",
				ID:            ".test.Fake.f_sint32",
				Typez:         SINT32_TYPE,
			},
			{
				Documentation: "A singular field tag = 18",
				Name:          "f_sint64",
				JSONName:      "fSint64",
				ID:            ".test.Fake.f_sint64",
				Typez:         SINT64_TYPE,
			},
		},
	})
}

func TestProtobuf_ScalarArray(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "scalar_array.proto"))
	message, ok := test.State.MessageByID[".test.Fake"]
	if !ok {
		t.Fatalf("Cannot find message %s in API State", ".test.Fake")
	}
	checkProtobufMessage(t, *message, Message{
		Name:          "Fake",
		Package:       "test",
		ID:            ".test.Fake",
		Documentation: "A test message.",
		Fields: []*Field{
			{
				Repeated:      true,
				Documentation: "A repeated field tag = 1",
				Name:          "f_double",
				JSONName:      "fDouble",
				ID:            ".test.Fake.f_double",
				Typez:         DOUBLE_TYPE,
			},
			{
				Repeated:      true,
				Documentation: "A repeated field tag = 3",
				Name:          "f_int64",
				JSONName:      "fInt64",
				ID:            ".test.Fake.f_int64",
				Typez:         INT64_TYPE,
			},
			{
				Repeated:      true,
				Documentation: "A repeated field tag = 9",
				Name:          "f_string",
				JSONName:      "fString",
				ID:            ".test.Fake.f_string",
				Typez:         STRING_TYPE,
			},
			{
				Repeated:      true,
				Documentation: "A repeated field tag = 12",
				Name:          "f_bytes",
				JSONName:      "fBytes",
				ID:            ".test.Fake.f_bytes",
				Typez:         BYTES_TYPE,
			},
		},
	})
}

func TestProtobuf_ScalarOptional(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "scalar_optional.proto"))
	message, ok := test.State.MessageByID[".test.Fake"]
	if !ok {
		t.Fatalf("Cannot find message %s in API", "Fake")
	}
	checkProtobufMessage(t, *message, Message{
		Name:          "Fake",
		Package:       "test",
		ID:            ".test.Fake",
		Documentation: "A test message.",
		Fields: []*Field{
			{
				Optional:      true,
				Documentation: "An optional field tag = 1",
				Name:          "f_double",
				JSONName:      "fDouble",
				ID:            ".test.Fake.f_double",
				Typez:         DOUBLE_TYPE,
			},
			{
				Optional:      true,
				Documentation: "An optional field tag = 3",
				Name:          "f_int64",
				JSONName:      "fInt64",
				ID:            ".test.Fake.f_int64",
				Typez:         INT64_TYPE,
			},
			{
				Optional:      true,
				Documentation: "An optional field tag = 9",
				Name:          "f_string",
				JSONName:      "fString",
				ID:            ".test.Fake.f_string",
				Typez:         STRING_TYPE,
			},
			{
				Optional:      true,
				Documentation: "An optional field tag = 12",
				Name:          "f_bytes",
				JSONName:      "fBytes",
				ID:            ".test.Fake.f_bytes",
				Typez:         BYTES_TYPE,
			},
		},
	})
}

func TestProtobuf_SkipExternalMessages(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "with_import.proto"))
	// Both `ImportedMessage` and `LocalMessage` should be in the index:
	_, ok := test.State.MessageByID[".away.ImportedMessage"]
	if !ok {
		t.Fatalf("Cannot find message %s in API State", ".away.ImportedMessage")
	}
	message, ok := test.State.MessageByID[".test.LocalMessage"]
	if !ok {
		t.Fatalf("Cannot find message %s in API State", ".test.LocalMessage")
	}
	checkProtobufMessage(t, *message, Message{
		Name:          "LocalMessage",
		Package:       "test",
		ID:            ".test.LocalMessage",
		Documentation: "This is a local message, it should be generated.",
		Fields: []*Field{
			{
				Name:          "payload",
				JSONName:      "payload",
				ID:            ".test.LocalMessage.payload",
				Documentation: "This field uses an imported message.",
				Typez:         MESSAGE_TYPE,
				TypezID:       ".away.ImportedMessage",
				Optional:      true,
			},
			{
				Name:          "value",
				JSONName:      "value",
				ID:            ".test.LocalMessage.value",
				Documentation: "This field uses an imported enum.",
				Typez:         ENUM_TYPE,
				TypezID:       ".away.ImportedEnum",
				Optional:      false,
			},
		},
	})
	// Only `LocalMessage` should be found in the messages list:
	for _, msg := range test.Messages {
		if msg.ID == ".test.ImportedMessage" {
			t.Errorf("imported messages should not be in message list %v", msg)
		}
	}
}

func TestProtobuf_SkipExternaEnums(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "with_import.proto"))
	// Both `ImportedEnum` and `LocalEnum` should be in the index:
	_, ok := test.State.EnumByID[".away.ImportedEnum"]
	if !ok {
		t.Fatalf("Cannot find enum %s in API State", ".away.ImportedEnum")
	}
	enum, ok := test.State.EnumByID[".test.LocalEnum"]
	if !ok {
		t.Fatalf("Cannot find enum %s in API State", ".test.LocalEnum")
	}
	checkProtobufEnum(t, *enum, Enum{
		Name:          "LocalEnum",
		Package:       "test",
		Documentation: "This is a local enum, it should be generated.",
		Values: []*EnumValue{
			{
				Name:   "RED",
				Number: 0,
			},
			{
				Name:   "WHITE",
				Number: 1,
			},
			{
				Name:   "BLUE",
				Number: 2,
			},
		},
	})
	// Only `LocalMessage` should be found in the messages list:
	for _, msg := range test.Messages {
		if msg.ID == ".test.ImportedMessage" {
			t.Errorf("imported messages should not be in message list %v", msg)
		}
	}
}

func TestProtobuf_Comments(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "comments.proto"))
	message, ok := test.State.MessageByID[".test.Request"]
	if !ok {
		t.Fatalf("Cannot find message %s in API State", ".test.Request")
	}
	checkProtobufMessage(t, *message, Message{
		Name:          "Request",
		Package:       "test",
		ID:            ".test.Request",
		Documentation: "A test message.\n\nWith even more of a description.\nMaybe in more than one line.\nAnd some markdown:\n- An item\n  - A nested item\n- Another item",
		Fields: []*Field{
			{
				Name:          "parent",
				Documentation: "A field.\n\nWith a longer description.",
				JSONName:      "parent",
				ID:            ".test.Request.parent",
				Typez:         STRING_TYPE,
			},
		},
	})

	message, ok = test.State.MessageByID[".test.Response.Nested"]
	if !ok {
		t.Fatalf("Cannot find message %s in API State", ".test.Response.nested")
	}
	checkProtobufMessage(t, *message, Message{
		Name:          "Nested",
		Package:       "test",
		ID:            ".test.Response.Nested",
		Documentation: "A nested message.\n\n- Item 1\n  Item 1 continued",
		Fields: []*Field{
			{
				Name:          "path",
				Documentation: "Field in a nested message.\n\n* Bullet 1\n  Bullet 1 continued\n* Bullet 2\n  Bullet 2 continued",
				JSONName:      "path",
				ID:            ".test.Response.Nested.path",
				Typez:         STRING_TYPE,
			},
		},
	})

	e, ok := test.State.EnumByID[".test.Response.Status"]
	if !ok {
		t.Fatalf("Cannot find enum %s in API State", ".test.Response.Status")
	}
	checkProtobufEnum(t, *e, Enum{
		Name:          "Status",
		Package:       "test",
		Documentation: "Some enum.\n\nLine 1.\nLine 2.",
		Values: []*EnumValue{
			{
				Name:          "NOT_READY",
				Documentation: "The first enum value description.\n\nValue Line 1.\nValue Line 2.",
				Number:        0,
			},
			{
				Name:          "READY",
				Documentation: "The second enum value description.",
				Number:        1,
			},
		},
	})

	service, ok := test.State.ServiceByID[".test.Service"]
	if !ok {
		t.Fatalf("Cannot find service %s in API State", ".test.Service")
	}
	checkProtobufService(t, *service, Service{
		Name:          "Service",
		ID:            ".test.Service",
		Package:       "test",
		Documentation: "A service.\n\nWith a longer service description.",
		DefaultHost:   "test.googleapis.com",
		Methods: []*Method{
			{
				Name:          "Create",
				ID:            ".test.Service.Create",
				Documentation: "Some RPC.\n\nIt does not do much.",
				InputTypeID:   ".test.Request",
				OutputTypeID:  ".test.Response",
				PathInfo: &PathInfo{
					Verb: "POST",
					PathTemplate: []PathSegment{
						NewLiteralPathSegment("v1"),
						NewFieldPathPathSegment("parent"),
						NewLiteralPathSegment("foos"),
					},
					QueryParameters: map[string]bool{},
					BodyFieldPath:   "*",
				},
			},
		},
	})
}

func TestProtobuf_OneOfs(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "oneofs.proto"))
	message, ok := test.State.MessageByID[".test.Fake"]
	if !ok {
		t.Fatalf("Cannot find message %s in API State", ".test.Request")
	}
	checkProtobufMessage(t, *message, Message{
		Name:          "Fake",
		Package:       "test",
		ID:            ".test.Fake",
		Documentation: "A test message.",
		Fields: []*Field{
			{
				Name:          "field_one",
				Documentation: "A string choice",
				JSONName:      "fieldOne",
				ID:            ".test.Fake.field_one",
				Typez:         STRING_TYPE,
				IsOneOf:       true,
			},
			{
				Documentation: "An int choice",
				Name:          "field_two",
				ID:            ".test.Fake.field_two",
				Typez:         INT64_TYPE,
				JSONName:      "fieldTwo",
				IsOneOf:       true,
			},
			{
				Documentation: "Optional is oneof in proto",
				Name:          "field_three",
				ID:            ".test.Fake.field_three",
				Typez:         STRING_TYPE,
				JSONName:      "fieldThree",
				Optional:      true,
			},
			{
				Documentation: "A normal field",
				Name:          "field_four",
				ID:            ".test.Fake.field_four",
				Typez:         INT32_TYPE,
				JSONName:      "fieldFour",
			},
		},
		OneOfs: []*OneOf{
			{
				Name: "choice",
				ID:   ".test.Fake.choice",
				Fields: []*Field{
					{
						Documentation: "A string choice",
						Name:          "field_one",
						ID:            ".test.Fake.field_one",
						Typez:         9,
						JSONName:      "fieldOne",
						IsOneOf:       true,
					},
					{
						Documentation: "An int choice",
						Name:          "field_two",
						ID:            ".test.Fake.field_two",
						Typez:         3,
						JSONName:      "fieldTwo",
						IsOneOf:       true,
					},
				},
			},
		},
	})
}

func TestProtobuf_ObjectFields(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "object_fields.proto"))
	message, ok := test.State.MessageByID[".test.Fake"]
	if !ok {
		t.Fatalf("Cannot find message %s in API State", ".test.Fake")
	}
	checkProtobufMessage(t, *message, Message{
		Name:    "Fake",
		Package: "test",
		ID:      ".test.Fake",
		Fields: []*Field{
			{
				Repeated: false,
				Optional: true,
				Name:     "singular_object",
				JSONName: "singularObject",
				ID:       ".test.Fake.singular_object",
				Typez:    MESSAGE_TYPE,
				TypezID:  ".test.Other",
			},
			{
				Repeated: true,
				Optional: false,
				Name:     "repeated_object",
				JSONName: "repeatedObject",
				ID:       ".test.Fake.repeated_object",
				Typez:    MESSAGE_TYPE,
				TypezID:  ".test.Other",
			},
		},
	})
}

func TestProtobuf_WellKnownTypeFields(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "wkt_fields.proto"))
	message, ok := test.State.MessageByID[".test.Fake"]
	if !ok {
		t.Fatalf("Cannot find message %s in API State", ".test.Fake")
	}
	checkProtobufMessage(t, *message, Message{
		Name:    "Fake",
		Package: "test",
		ID:      ".test.Fake",
		Fields: []*Field{
			{
				Name:     "field_mask",
				JSONName: "fieldMask",
				ID:       ".test.Fake.field_mask",
				Typez:    MESSAGE_TYPE,
				TypezID:  ".google.protobuf.FieldMask",
				Optional: true,
			},
			{
				Name:     "timestamp",
				JSONName: "timestamp",
				ID:       ".test.Fake.timestamp",
				Typez:    MESSAGE_TYPE,
				TypezID:  ".google.protobuf.Timestamp",
				Optional: true,
			},
			{
				Name:     "any",
				JSONName: "any",
				ID:       ".test.Fake.any",
				Typez:    MESSAGE_TYPE,
				TypezID:  ".google.protobuf.Any",
				Optional: true,
			},
			{
				Name:     "repeated_field_mask",
				JSONName: "repeatedFieldMask",
				ID:       ".test.Fake.repeated_field_mask",
				Typez:    MESSAGE_TYPE,
				TypezID:  ".google.protobuf.FieldMask",
				Repeated: true,
			},
			{
				Name:     "repeated_timestamp",
				JSONName: "repeatedTimestamp",
				ID:       ".test.Fake.repeated_timestamp",
				Typez:    MESSAGE_TYPE,
				TypezID:  ".google.protobuf.Timestamp",
				Repeated: true,
			},
			{
				Name:     "repeated_any",
				JSONName: "repeatedAny",
				ID:       ".test.Fake.repeated_any",
				Typez:    MESSAGE_TYPE,
				TypezID:  ".google.protobuf.Any",
				Repeated: true,
			},
		},
	})
}

func TestProtobuf_MapFields(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "map_fields.proto"))
	message, ok := test.State.MessageByID[".test.Fake"]
	if !ok {
		t.Fatalf("Cannot find message %s in API State", ".test.Fake")
	}
	checkProtobufMessage(t, *message, Message{
		Name:    "Fake",
		Package: "test",
		ID:      ".test.Fake",
		Fields: []*Field{
			{
				Repeated: false,
				Optional: false,
				Name:     "singular_map",
				JSONName: "singularMap",
				ID:       ".test.Fake.singular_map",
				Typez:    MESSAGE_TYPE,
				TypezID:  ".test.Fake.SingularMapEntry",
			},
		},
	})

	message, ok = test.State.MessageByID[".test.Fake.SingularMapEntry"]
	if !ok {
		t.Fatalf("Cannot find message %s in API State", ".test.Fake")
	}
	checkProtobufMessage(t, *message, Message{
		Name:    "SingularMapEntry",
		Package: "test",
		ID:      ".test.Fake.SingularMapEntry",
		IsMap:   true,
		Fields: []*Field{
			{
				Repeated: false,
				Optional: false,
				Name:     "key",
				JSONName: "key",
				ID:       ".test.Fake.SingularMapEntry.key",
				Typez:    STRING_TYPE,
			},
			{
				Repeated: false,
				Optional: false,
				Name:     "value",
				JSONName: "value",
				ID:       ".test.Fake.SingularMapEntry.value",
				Typez:    INT32_TYPE,
			},
		},
	})
}

func TestProtobuf_Service(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "test_service.proto"))
	service, ok := test.State.ServiceByID[".test.TestService"]
	if !ok {
		t.Fatalf("Cannot find service %s in API State", ".test.TestService")
	}
	checkProtobufService(t, *service, Service{
		Name:          "TestService",
		Package:       "test",
		ID:            ".test.TestService",
		Documentation: "A service to unit test the protobuf translator.",
		DefaultHost:   "test.googleapis.com",
		Methods: []*Method{
			{
				Name:          "GetFoo",
				ID:            ".test.TestService.GetFoo",
				Documentation: "Gets a Foo resource.",
				InputTypeID:   ".test.GetFooRequest",
				OutputTypeID:  ".test.Foo",
				PathInfo: &PathInfo{
					Verb: "GET",
					PathTemplate: []PathSegment{
						NewLiteralPathSegment("v1"),
						NewFieldPathPathSegment("name"),
					},
					QueryParameters: map[string]bool{},
					BodyFieldPath:   "",
				},
			},
			{
				Name:          "CreateFoo",
				ID:            ".test.TestService.CreateFoo",
				Documentation: "Creates a new Foo resource.",
				InputTypeID:   ".test.CreateFooRequest",
				OutputTypeID:  ".test.Foo",
				PathInfo: &PathInfo{
					Verb: "POST",
					PathTemplate: []PathSegment{
						NewLiteralPathSegment("v1"),
						NewFieldPathPathSegment("parent"),
						NewLiteralPathSegment("foos"),
					},
					QueryParameters: map[string]bool{"foo_id": true},
					BodyFieldPath:   "foo",
				},
			},
		},
	})
}

func TestProtobuf_QueryParameters(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "query_parameters.proto"))
	service, ok := test.State.ServiceByID[".test.TestService"]
	if !ok {
		t.Fatalf("Cannot find service %s in API State", ".test.TestService")
	}
	checkProtobufService(t, *service, Service{
		Name:          "TestService",
		Package:       "test",
		ID:            ".test.TestService",
		Documentation: "A service to unit test the protobuf translator.",
		DefaultHost:   "test.googleapis.com",
		Methods: []*Method{
			{
				Name:          "CreateFoo",
				ID:            ".test.TestService.CreateFoo",
				Documentation: "Creates a new `Foo` resource. `Foo`s are containers for `Bar`s.\n\nShows how a `body: \"${field}\"` option works.",
				InputTypeID:   ".test.CreateFooRequest",
				OutputTypeID:  ".test.Foo",
				PathInfo: &PathInfo{
					Verb: "POST",
					PathTemplate: []PathSegment{
						NewLiteralPathSegment("v1"),
						NewFieldPathPathSegment("parent"),
						NewLiteralPathSegment("foos"),
					},
					QueryParameters: map[string]bool{"foo_id": true},
					BodyFieldPath:   "bar",
				},
			},
			{
				Name:          "AddBar",
				ID:            ".test.TestService.AddBar",
				Documentation: "Add a Bar resource.\n\nShows how a `body: \"*\"` option works.",
				InputTypeID:   ".test.AddBarRequest",
				OutputTypeID:  ".test.Bar",
				PathInfo: &PathInfo{
					Verb: "POST",
					PathTemplate: []PathSegment{
						NewLiteralPathSegment("v1"),
						NewFieldPathPathSegment("parent"),
						NewVerbPathSegment("addFoo"),
					},
					QueryParameters: map[string]bool{},
					BodyFieldPath:   "*",
				},
			},
		},
	})
}

func TestProtobuf_Enum(t *testing.T) {
	test := makeAPIForProtobuf(nil, newTestCodeGeneratorRequest(t, "enum.proto"))
	e, ok := test.State.EnumByID[".test.Code"]
	if !ok {
		t.Fatalf("Cannot find enum %s in API State", ".test.Code")
	}
	checkProtobufEnum(t, *e, Enum{
		Name:          "Code",
		Package:       "test",
		Documentation: "An enum.",
		Values: []*EnumValue{
			{
				Name:          "OK",
				Documentation: "Not an error; returned on success.",
				Number:        0,
			},
			{
				Name:          "UNKNOWN",
				Documentation: "Unknown error.",
				Number:        1,
			},
		},
	})
}

func TestProtobuf_TrimLeadingSpacesInDocumentation(t *testing.T) {
	input := ` In this example, in proto field could take one of the following values:

 * full_name for a violation in the full_name value
 * email_addresses[1].email for a violation in the email field of the
   first email_addresses message
 * email_addresses[3].type[2] for a violation in the second type
   value in the third email_addresses message.)`

	want := `In this example, in proto field could take one of the following values:

* full_name for a violation in the full_name value
* email_addresses[1].email for a violation in the email field of the
  first email_addresses message
* email_addresses[3].type[2] for a violation in the second type
  value in the third email_addresses message.)`

	got := trimLeadingSpacesInDocumentation(input)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("mismatch in FormatDocComments (-want, +got)\n:%s", diff)
	}
}

func TestProtobuf_LocationMixin(t *testing.T) {
	var serviceConfig = &serviceconfig.Service{
		Name:  "test.googleapis.com",
		Title: "Test API",
		Documentation: &serviceconfig.Documentation{
			Summary:  "Used for testing generation.",
			Overview: "Test Overview",
		},
		Apis: []*apipb.Api{
			{
				Name: "google.cloud.location.Locations",
			},
			{
				Name: "test.googleapis.com.TestService",
			},
		},
		Http: &annotations.Http{
			Rules: []*annotations.HttpRule{
				{
					Selector: "google.cloud.location.Locations.GetLocation",
					Pattern: &annotations.HttpRule_Get{
						Get: "/v1/{name=projects/*/locations/*}",
					},
				},
			},
		},
	}
	test := makeAPIForProtobuf(serviceConfig, newTestCodeGeneratorRequest(t, "test_service.proto"))
	service, ok := test.State.ServiceByID[".google.cloud.location.Locations"]
	if !ok {
		t.Fatalf("Cannot find service %s in API State", ".google.cloud.location.Locations")
	}
	checkProtobufService(t, *service, Service{
		Documentation: "Manages location-related information with an API service.",
		DefaultHost:   "cloud.googleapis.com",
		Name:          "Locations",
		ID:            ".google.cloud.location.Locations",
		Package:       "google.cloud.location",
		Methods: []*Method{
			{
				Documentation: "GetLocation is an RPC method of Locations.",
				Name:          "GetLocation",
				ID:            ".google.cloud.location.Locations.GetLocation",
				InputTypeID:   ".google.cloud.location.GetLocationRequest",
				OutputTypeID:  ".google.cloud.location.Location",
				PathInfo: &PathInfo{
					Verb: "GET",
					PathTemplate: []PathSegment{
						NewLiteralPathSegment("v1"),
						NewFieldPathPathSegment("name"),
					},
					QueryParameters: map[string]bool{},
				},
			},
		},
	})
}

func TestProtobuf_IAMMixin(t *testing.T) {
	var serviceConfig = &serviceconfig.Service{
		Name:  "test.googleapis.com",
		Title: "Test API",
		Documentation: &serviceconfig.Documentation{
			Summary:  "Used for testing generation.",
			Overview: "Test Overview",
		},
		Apis: []*apipb.Api{
			{
				Name: "google.iam.v1.IAMPolicy",
			},
			{
				Name: "test.googleapis.com.TestService",
			},
		},
		Http: &annotations.Http{
			Rules: []*annotations.HttpRule{
				{
					Selector: "google.iam.v1.IAMPolicy.GetIamPolicy",
					Pattern: &annotations.HttpRule_Post{
						Post: "/v1/{resource=services/*}:getIamPolicy",
					},
					Body: "*",
				},
			},
		},
	}
	test := makeAPIForProtobuf(serviceConfig, newTestCodeGeneratorRequest(t, "test_service.proto"))
	service, ok := test.State.ServiceByID[".google.iam.v1.IAMPolicy"]
	if !ok {
		t.Fatalf("Cannot find service %s in API State", ".google.iam.v1.IAMPolicy")
	}
	checkProtobufService(t, *service, Service{
		Documentation: "Manages Identity and Access Management (IAM) policies with an API service.",
		DefaultHost:   "iam-meta-api.googleapis.com",
		Name:          "IAMPolicy",
		ID:            ".google.iam.v1.IAMPolicy",
		Package:       "google.iam.v1",
		Methods: []*Method{
			{
				Documentation: "GetIamPolicy is an RPC method of IAMPolicy.",
				Name:          "GetIamPolicy",
				ID:            ".google.iam.v1.IAMPolicy.GetIamPolicy",
				InputTypeID:   ".google.iam.v1.GetIamPolicyRequest",
				OutputTypeID:  ".google.iam.v1.Policy",
				PathInfo: &PathInfo{
					Verb: "POST",
					PathTemplate: []PathSegment{
						NewLiteralPathSegment("v1"),
						NewFieldPathPathSegment("resource"),
						NewVerbPathSegment("getIamPolicy"),
					},
					QueryParameters: map[string]bool{},
					BodyFieldPath:   "*",
				},
			},
		},
	})
}

func TestProtobuf_OperationMixin(t *testing.T) {
	var serviceConfig = &serviceconfig.Service{
		Name:  "test.googleapis.com",
		Title: "Test API",
		Documentation: &serviceconfig.Documentation{
			Summary:  "Used for testing generation.",
			Overview: "Test Overview",
			Rules: []*serviceconfig.DocumentationRule{
				{
					Selector:    "google.longrunning.Operations.GetOperation",
					Description: "Custom docs.",
				},
			},
		},
		Apis: []*apipb.Api{
			{
				Name: "google.longrunning.Operations",
			},
			{
				Name: "test.googleapis.com.TestService",
			},
		},
		Http: &annotations.Http{
			Rules: []*annotations.HttpRule{
				{
					Selector: "google.longrunning.Operations.GetOperation",
					Pattern: &annotations.HttpRule_Get{
						Get: "/v2/{name=operations/*}",
					},
					Body: "*",
				},
			},
		},
	}
	test := makeAPIForProtobuf(serviceConfig, newTestCodeGeneratorRequest(t, "test_service.proto"))
	service, ok := test.State.ServiceByID[".google.longrunning.Operations"]
	if !ok {
		t.Fatalf("Cannot find service %s in API State", ".google.longrunning.Operations")
	}
	checkProtobufService(t, *service, Service{
		Documentation: "Manages long-running operations with an API service.",
		DefaultHost:   "longrunning.googleapis.com",
		Name:          "Operations",
		ID:            ".google.longrunning.Operations",
		Package:       "google.longrunning",
		Methods: []*Method{
			{
				Documentation: "Custom docs.",
				Name:          "GetOperation",
				ID:            ".google.longrunning.Operations.GetOperation",
				InputTypeID:   ".google.longrunning.GetOperationRequest",
				OutputTypeID:  ".google.longrunning.Operation",
				PathInfo: &PathInfo{
					Verb: "GET",
					PathTemplate: []PathSegment{
						NewLiteralPathSegment("v2"),
						NewFieldPathPathSegment("name"),
					},
					QueryParameters: map[string]bool{},
					BodyFieldPath:   "*",
				},
			},
		},
	})
}

func newTestCodeGeneratorRequest(t *testing.T, filename string) *pluginpb.CodeGeneratorRequest {
	t.Helper()
	options := map[string]string{
		"googleapis-root": "../testing/testdata/googleapis",
		"test-root":       "testdata",
	}
	request, err := newCodeGeneratorRequest(filename, options)
	if err != nil {
		t.Fatal(err)
	}
	return request
}

func checkProtobufMessage(t *testing.T, got Message, want Message) {
	t.Helper()
	// Checking Parent, Messages, Fields, and OneOfs requires special handling.
	if diff := cmp.Diff(want, got, cmpopts.IgnoreFields(Message{}, "Fields", "OneOfs", "Parent", "Messages")); diff != "" {
		t.Errorf("message attributes mismatch (-want +got):\n%s", diff)
	}
	less := func(a, b *Field) bool { return a.Name < b.Name }
	if diff := cmp.Diff(want.Fields, got.Fields, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("field mismatch (-want, +got):\n%s", diff)
	}
	// Ignore parent because types are cyclic
	if diff := cmp.Diff(want.OneOfs, got.OneOfs, cmpopts.SortSlices(less), cmpopts.IgnoreFields(OneOf{}, "Parent")); diff != "" {
		t.Errorf("oneofs mismatch (-want, +got):\n%s", diff)
	}
}

func checkProtobufEnum(t *testing.T, got Enum, want Enum) {
	t.Helper()
	if diff := cmp.Diff(want, got, cmpopts.IgnoreFields(Enum{}, "Values", "Parent")); diff != "" {
		t.Errorf("Mismatched service attributes (-want, +got):\n%s", diff)
	}
	less := func(a, b *EnumValue) bool { return a.Name < b.Name }
	if diff := cmp.Diff(want.Values, got.Values, cmpopts.SortSlices(less), cmpopts.IgnoreFields(EnumValue{}, "Parent")); diff != "" {
		t.Errorf("method mismatch (-want, +got):\n%s", diff)
	}
}

func checkProtobufService(t *testing.T, got Service, want Service) {
	t.Helper()
	if diff := cmp.Diff(want, got, cmpopts.IgnoreFields(Service{}, "Methods")); diff != "" {
		t.Errorf("Mismatched service attributes (-want, +got):\n%s", diff)
	}
	less := func(a, b *Method) bool { return a.Name < b.Name }
	if diff := cmp.Diff(want.Methods, got.Methods, cmpopts.SortSlices(less)); diff != "" {
		t.Errorf("method mismatch (-want, +got):\n%s", diff)
	}
}