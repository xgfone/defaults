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
	"net"
	"net/http"
)

var (
	// GetRemoteAddrFunc is used to get the remote address.
	//
	// For the default implementation, it only detects req
	// and supports the types or interfaces:
	//
	// 	*http.Request
	//	interface{ RemoteAddr() string }
	//	interface{ RemoteAddr() net.Addr }
	//
	// Or, return "".
	GetRemoteAddrFunc = NewValueWithValidation(getRemoteAddr, fActxAifaceR1[string]("GetRemoteAddr"))
)

// GetRemoteAddr is the proxy of GetRemoteAddrFunc to call the function.
func GetRemoteAddr(ctx context.Context, req interface{}) string {
	return GetRemoteAddrFunc.Get()(ctx, req)
}

func getRemoteAddr(ctx context.Context, req interface{}) string {
	switch v := req.(type) {
	case interface{ RemoteAddr() string }:
		return v.RemoteAddr()

	case interface{ RemoteAddr() net.Addr }:
		return v.RemoteAddr().String()

	case *http.Request:
		return v.RemoteAddr

	default:
		return ""
	}
}
