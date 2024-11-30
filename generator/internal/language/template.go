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

type TemplateData struct {
	TemplateDir      string
	Name             string
	Title            string
	Description      string
	PackageName      string
	RequiredPackages []string
	HasServices      bool
	Imports          []string
	DefaultHost      string
	Services         []*Service
	Messages         []*Message
	NameToLower      string
}

type Service struct {
	Methods             []*Method
	NameToSnake         string
	NameToPascal        string
	ServiceNameToPascal string
	NameToCamel         string
	ServiceName         string
	DocLines            []string
	DefaultHost         string
}

type Message struct {
	Fields            []*Field
	BasicFields       []*Field
	ExplicitOneOfs    []*OneOf
	NestedMessages    []*Message
	Enums             []*Enum
	MessageAttributes []string
	Name              string
	QualifiedName     string
	NameSnakeCase     string
	HasNestedTypes    bool
	DocLines          []string
	IsMap             bool
}

type Method struct {
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
	QueryParams       []*Field
	HasBody           bool
	BodyAccessor      string
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
