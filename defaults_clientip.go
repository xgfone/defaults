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
	"net/netip"

	"github.com/xgfone/go-defaults/assists"
)

var (
	// GetClientIPFunc is used to get the client ip of the request.
	//
	// For the default implementation, it only supports the types or interfaces:
	//
	// 	*http.Request
	// 	interface{ ClientIP() netip.Addr }
	// 	interface{ ClientIP() net.IP }
	// 	interface{ ClientIP() string }
	// 	interface{ RemoteAddr() netip.Addr }
	// 	interface{ RemoteAddr() net.Addr }
	// 	interface{ RemoteAddr() string }
	//
	// Or, return nil.
	GetClientIPFunc = NewValueWithValidation(getClientIP, fActxAifaceR1[netip.Addr]("GetClientIP"))
)

// GetClientIP is the proxy of GetClientIPFunc to call the function.
func GetClientIP(ctx context.Context, req interface{}) netip.Addr {
	return GetClientIPFunc.Get()(ctx, req)
}

func getClientIP(ctx context.Context, req interface{}) netip.Addr {
	switch v := req.(type) {
	case interface{ ClientIP() netip.Addr }:
		return v.ClientIP()

	case interface{ ClientIP() net.IP }:
		return ip2addr(v.ClientIP())

	case interface{ ClientIP() string }:
		return str2addr(v.ClientIP())

	case interface{ RemoteAddr() netip.Addr }:
		return v.RemoteAddr()

	case interface{ RemoteAddr() net.Addr }:
		return assists.ConvertAddr(v.RemoteAddr())

	case interface{ RemoteAddr() string }:
		return str2addr(v.RemoteAddr())

	case *http.Request:
		return str2addr(v.RemoteAddr)

	default:
		return netip.Addr{}
	}
}

func ip2addr(ip net.IP) netip.Addr {
	addr, _ := netip.AddrFromSlice(ip)
	return addr
}

func str2addr(s string) netip.Addr {
	addr, _ := netip.ParseAddr(assists.TrimIP(s))
	return addr
}
