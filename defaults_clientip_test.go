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
	"context"
	"net/http"
	"testing"
)

func TestString2netip(t *testing.T) {
	expect := "127.0.0.1"
	result := string2netip("127.0.0.1:80").String()
	if result != expect {
		t.Errorf("expect '%s', but got '%s'", expect, result)
	}
}

func TestGetClientIP(t *testing.T) {
	expect := "127.0.0.1"
	ip := GetClientIP(context.Background(), &http.Request{RemoteAddr: "127.0.0.1:80"})
	if result := ip.String(); result != expect {
		t.Errorf("expect '%s', but got '%s'", expect, result)
	}
}
