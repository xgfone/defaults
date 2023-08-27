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

package assists

import "testing"

func TestString2netip(t *testing.T) {
	expect := "127.0.0.1"
	result := str2addr("127.0.0.1:80").String()
	if result != expect {
		t.Errorf("expect '%s', but got '%s'", expect, result)
	}

	expect = "ff00::"
	result = str2addr("[ff00::]:80").String()
	if result != expect {
		t.Errorf("expect '%s', but got '%s'", expect, result)
	}
}
