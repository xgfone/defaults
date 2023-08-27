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
	"net/netip"
	"testing"
)

func BenchmarkGetClientIP(b *testing.B) {
	var drop func(netip.Addr)

	r := &http.Request{RemoteAddr: "127.0.0.1:80"}
	drop = func(netip.Addr) {}

	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			ip := GetClientIP(context.Background(), r)
			drop(ip)
		}
	})
}

func TestString2netip(t *testing.T) {
	expect := "127.0.0.1"
	result := str2addr("127.0.0.1:80").String()
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
