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
	// GetClientIPFunc is used to get the client ip of the request.
	//
	// For the default implementation, it only supports the types or interfaces:
	//
	// 	*http.Request
	// 	interface{ ClientIP() net.IP }
	// 	interface{ ClientIP() string }
	// 	interface{ RemoteAddr() net.Addr }
	// 	interface{ RemoteAddr() string }
	//
	// Or, return nil.
	GetClientIPFunc = NewValueWithValidation(getClientIP, fActxAifaceR1[net.IP]("GetClientIP"))
)

// GetClientIP is the proxy of GetClientIPFunc to call the function.
func GetClientIP(ctx context.Context, req interface{}) net.IP {
	return GetClientIPFunc.Get()(ctx, req)
}

func getClientIP(ctx context.Context, req interface{}) net.IP {
	switch v := req.(type) {
	case interface{ ClientIP() net.IP }:
		return v.ClientIP()

	case interface{ ClientIP() string }:
		return net.ParseIP(v.ClientIP())

	case interface{ RemoteAddr() net.Addr }:
		return netaddr2netip(v.RemoteAddr())

	case interface{ RemoteAddr() string }:
		return string2netip(v.RemoteAddr())

	case *http.Request:
		return string2netip(v.RemoteAddr)

	default:
		return nil
	}
}

func netaddr2netip(addr net.Addr) net.IP {
	switch v := addr.(type) {
	case *net.TCPAddr:
		return v.IP

	case *net.UDPAddr:
		return v.IP

	default:
		return string2netip(v.String())
	}
}

func string2netip(s string) net.IP {
	if host, _, _ := net.SplitHostPort(s); host != "" {
		s = host
	}
	return net.ParseIP(s)
}
