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

// TrimIP trims the port from the string s if exists. Or, return s itself.
//
// NOTICE: s must be a valid format "IP: or "IP:PORT".
func TrimIP(s string) string {
	for i, _len := 0, len(s); i < _len; i++ {
		switch s[i] {
		case '.': // IPv4, such as "1.2.3.4" or "1.2.3.4:80"
			if index := strings.LastIndexByte(s, ':'); index > 0 {
				s = s[:index]
			}
			return s

		case ':', '[': // IPv6, such as "ff00::" or "[ff00::]:80"
			if index := strings.LastIndexByte(s, ']'); index > 0 {
				s = s[1:index]
			}
			return s
		}
	}
	return s
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
		addr, _ = netip.ParseAddr(TrimIP(v.String()))
	}
	return
}

// IP2Addr converts net.IP to netip.Addr, which also converts 4in6 to ipv4.
//
// If ip is invalid, the returned addr is also invalid.
func IP2Addr(ip net.IP) (addr netip.Addr) {
	switch len(ip) {
	case net.IPv4len:
		addr = netip.AddrFrom4(*(*[4]byte)(ip))
		// addr = netip.AddrFrom4([4]byte(ip)) // 1.20+
	case net.IPv6len:
		if ipv4 := ip.To4(); ipv4 != nil {
			addr = netip.AddrFrom4(*(*[4]byte)(ipv4))
			// addr = netip.AddrFrom4([4]byte(ipv4)) // 1.20+
		} else {
			addr = netip.AddrFrom16(*(*[16]byte)(ip))
			// addr = netip.AddrFrom16([16]byte(ip)) // 1.20+
		}
	}
	return
}
