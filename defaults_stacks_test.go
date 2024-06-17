// Copyright 2023 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package defaults

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetStacks(t *testing.T) {
	stacks := GetStacks(0)
	for i, stack := range stacks {
		if strings.HasPrefix(stack, "testing/") {
			stacks = stacks[:i]
			break
		}
	}

	expects := []string{
		"github.com/xgfone/go-defaults/defaults_stacks_test.go:TestGetStacks:24",
	}

	if len(expects) != len(stacks) {
		t.Fatalf("expect %d line, but got %d: %v", len(expects), len(stacks), stacks)
	}

	for i, line := range expects {
		if line != stacks[i] {
			t.Errorf("%d: expect '%s', but got '%s'", i, line, stacks[i])
		}
	}
}

func ExampleTrimPkgFile() {
	srcfile := TrimPkgFile("/path/to/src/github.com/xgfone/go-defaults/srcfile.go")
	modfile := TrimPkgFile("/path/to/pkg/mod/github.com/xgfone/go-defaults/modfile.go")
	origfile := TrimPkgFile("/path/to/github.com/xgfone/go-defaults/modfile.go")

	fmt.Println(srcfile)
	fmt.Println(modfile)
	fmt.Println(origfile)

	// Output:
	// github.com/xgfone/go-defaults/srcfile.go
	// github.com/xgfone/go-defaults/modfile.go
	// /path/to/github.com/xgfone/go-defaults/modfile.go
}

func TestGetCaller(t *testing.T) {
	caller := GetCaller(0)
	expect := "github.com/xgfone/go-defaults/defaults_stacks_test.go:TestGetCaller:63"
	if caller != expect {
		t.Errorf("expect '%s', but got '%s'", expect, caller)
	}
}
