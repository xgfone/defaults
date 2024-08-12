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
	"net"
	"net/netip"
	"strings"
)

// TrimPort trims the port from the host if exists. Or, returns itself.
//
// host may be one of formats as follow:
//
//	"ipv4"
//	"ipv6"
//	"domain"
//	"ipv4:port"
//	"[ipv6]:port"
//	"domain:port"
//
// Any that is not an ipv4 or ipv6 is as a domain.
//
// NOTICE: it does not validate whether the host is valid.
func TrimPort(host string) string {
	if host == "" {
		return ""
	}

	// For "[IPv6]:Port"
	if host[0] == '[' {
		if index := strings.LastIndexByte(host, ']'); index > 0 {
			host = host[1:index]
		}
		return host
	}

	// For "IPv4" or "Domain"
	lastindex := strings.LastIndexByte(host, ':')
	if lastindex == -1 {
		return host
	}

	// For "IPv4:Port", or "Domain:Port"
	if strings.IndexByte(host, ':') == lastindex {
		return host[:lastindex]
	}

	// Default For "IPv6"
	return host
}

// ConvertAddr converts the address from net.Addr to netip.Addr.
//
// If failed, return ZERO.
func ConvertAddr(netaddr net.Addr) (addr netip.Addr) {
	switch v := netaddr.(type) {
	case *net.TCPAddr:
		addr = IP2Addr(v.IP)

	case *net.UDPAddr:
		addr = IP2Addr(v.IP)

	default:
		addr, _ = netip.ParseAddr(TrimPort(v.String()))
	}
	return
}

// IP2Addr converts net.IP to netip.Addr, which also converts 4in6 to ipv4.
//
// If ip is invalid, the returned addr is also invalid.
func IP2Addr(ip net.IP) (addr netip.Addr) {
	switch len(ip) {
	case net.IPv4len:
		addr = netip.AddrFrom4([4]byte(ip))
	case net.IPv6len:
		if ipv4 := ip.To4(); ipv4 != nil {
			addr = netip.AddrFrom4([4]byte(ipv4))
		} else {
			addr = netip.AddrFrom16([16]byte(ip))
		}
	}
	return
}
