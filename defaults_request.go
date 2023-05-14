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
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/netip"
)

var (
	// HeaderXRequestID is used by GetRequestIDFunc to try
	// to get the request id from the http request.
	HeaderXRequestID = "X-Request-Id"

	// GetRequestIDFunc is used to get the unique request session id.
	//
	// For the default implementation, it only detects req
	// and supports the types or interfaces:
	//
	//	*http.Request
	//	interface{ RequestID() string }
	//	interface{ GetRequestID() string }
	//  interface{ Request() *http.Request }
	//	interface{ GetRequest() *http.Request }
	//  interface{ HTTPRequest() *http.Request }
	//	interface{ GetHTTPRequest() *http.Request }
	//
	// For *http.Request, it will returns the header HeaderXRequestID.
	//
	// Return "" instead if not found.
	GetRequestIDFunc = NewValueWithValidation(getRequestID, reqidValidateFunc)

	// GetRemoteAddrFunc is used to get the remote address.
	//
	// For the default implementation, it only detects req
	// and supports the types or interfaces:
	//
	//	*http.Request
	//  interface{ Request() *http.Request }
	//	interface{ GetRequest() *http.Request }
	//  interface{ HTTPRequest() *http.Request }
	//	interface{ GetHTTPRequest() *http.Request }
	//	interface{ RemoteAddr() string }
	//	interface{ RemoteAddr() net.Addr }
	//	interface{ RemoteAddr() netip.AddrPort }
	GetRemoteAddrFunc = NewValueWithValidation(getRemoteAddr, raddrValidateFunc)
)

// GetRequestID is the proxy of GetRequestIDFunc to call the function.
func GetRequestID(ctx context.Context, req interface{}) string {
	return GetRequestIDFunc.Get()(ctx, req)
}

// GetRemoteAddr is the proxy of GetRemoteAddrFunc to call the function.
func GetRemoteAddr(ctx context.Context, req interface{}) string {
	return GetRemoteAddrFunc.Get()(ctx, req)
}

func reqidValidateFunc(f func(context.Context, interface{}) string) error {
	if f == nil {
		return errors.New("GetRequestID function must not be nil")
	}
	return nil
}

func raddrValidateFunc(f func(context.Context, interface{}) string) error {
	if f == nil {
		return errors.New("GetRemoteAddr function must not be nil")
	}
	return nil
}

func getRequestID(ctx context.Context, req interface{}) string {
	switch r := req.(type) {
	case *http.Request:
		return r.Header.Get(HeaderXRequestID)

	case interface{ RequestID() string }:
		return r.RequestID()

	case interface{ GetRequestID() string }:
		return r.GetRequestID()

	case interface{ Request() *http.Request }:
		return r.Request().Header.Get(HeaderXRequestID)

	case interface{ HTTPRequest() *http.Request }:
		return r.HTTPRequest().Header.Get(HeaderXRequestID)

	case interface{ GetRequest() *http.Request }:
		return r.GetRequest().Header.Get(HeaderXRequestID)

	case interface{ GetHTTPRequest() *http.Request }:
		return r.GetHTTPRequest().Header.Get(HeaderXRequestID)

	default:
		return ""
	}
}

func getRemoteAddr(ctx context.Context, req interface{}) string {
	switch v := req.(type) {
	case *http.Request:
		return v.RemoteAddr

	case interface{ RemoteAddr() string }:
		return v.RemoteAddr()

	case interface{ RemoteAddr() net.Addr }:
		return v.RemoteAddr().String()

	case interface{ RemoteAddr() netip.AddrPort }:
		return v.RemoteAddr().String()

	case interface{ Request() *http.Request }:
		return v.Request().RemoteAddr

	case interface{ HTTPRequest() *http.Request }:
		return v.HTTPRequest().RemoteAddr

	case interface{ GetRequest() *http.Request }:
		return v.GetRequest().RemoteAddr

	case interface{ GetHTTPRequest() *http.Request }:
		return v.GetHTTPRequest().RemoteAddr

	default:
		panic(fmt.Errorf("GetRemoteAddr: unknown type %T", req))
	}
}
