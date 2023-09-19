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

import (
	"fmt"
	"net"
	"testing"
)

func ExampleTrimPort() {
	fmt.Println(TrimPort(""))
	fmt.Println(TrimPort("[abc"))
	fmt.Println(TrimPort("[abc]"))
	fmt.Println(TrimPort("[abc]:80")) // We just trim it by the lexical rule, not check its validity.
	fmt.Println(TrimPort("[ff00::]:80"))
	fmt.Println(TrimPort("ff00::"))
	fmt.Println(TrimPort("1.2.3.4"))
	fmt.Println(TrimPort("1.2.3.4:80"))
	fmt.Println(TrimPort("localhost"))
	fmt.Println(TrimPort("localhost:80"))

	// Output:
	//
	// [abc
	// abc
	// abc
	// ff00::
	// ff00::
	// 1.2.3.4
	// 1.2.3.4
	// localhost
	// localhost
}

func TestIP2Addr(t *testing.T) {
	if IP2Addr(nil).IsValid() {
		t.Error("expect invalid, but got valid")
	}

	if !IP2Addr(net.ParseIP("1.2.3.4")).Is4() {
		t.Error("expect an ipv4, but got not")
	}

	if !IP2Addr(net.ParseIP("ff00::")).Is6() {
		t.Error("expect an ipv4, but got not")
	}

	if !IP2Addr(net.ParseIP("::ffff:1.2.3.4")).Is4() {
		t.Error("expect an ipv4, but got not")
	}
}
