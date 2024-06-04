// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
)

// stringListVal implements the flag.Value interface and supports passing key
// value pairs multiple times on the command line.
type stringListVal []string

func newStringListVal(m *[]string) *stringListVal {
	return (*stringListVal)(m)
}

func (s *stringListVal) Set(val string) error {
	*s = append(*s, val)
	return nil
}

func (s *stringListVal) String() string {
	var b bytes.Buffer
	for i, key := range *s {
		if i > 0 {
			b.WriteRune(',')
		}
		b.WriteString(key)
	}
	return b.String()
}
