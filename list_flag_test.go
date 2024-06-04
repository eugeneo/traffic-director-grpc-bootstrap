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
	"flag"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestStringListVal(t *testing.T) {
	tests := []struct {
		desc      string
		keyValues []string
		wantList  []string
	}{
		{
			desc:      "happy single",
			keyValues: []string{"val"},
			wantList:  []string{"val"},
		},
		{
			desc:      "happy multiple",
			keyValues: []string{"val1", "val2"},
			wantList:  []string{"val1", "val2"},
		},
		{
			desc:      "happy with = in val",
			keyValues: []string{"val=1"},
			wantList:  []string{"val=1"},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			sm := []string{}
			fs := flag.NewFlagSet("testStringListVal", flag.ContinueOnError)
			fs.SetOutput(ioutil.Discard)
			fs.Var(newStringListVal(&sm), "metadata", "")

			var cmdLine []string
			for _, kv := range test.keyValues {
				cmdLine = append(cmdLine, "-metadata", kv)
			}
			if err := fs.Parse(cmdLine); err != nil {
				t.Fatalf("Parse(%v) returned err: %v", cmdLine, err)
			}
			if !cmp.Equal(sm, test.wantList, cmpopts.EquateEmpty()) {
				t.Fatalf("stringList after Parse(%v) is: %v, want: %v", cmdLine, sm, test.wantList)
			}
		})
	}
}
