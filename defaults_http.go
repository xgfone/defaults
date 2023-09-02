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
)

var (
	// HTTPIsRespondedFunc is used to report whether the http response is responded.
	//
	// For the default implementation, it will try check http.ResponseWriter:
	//   1. Check whether it implements the interface{ WroteHeader() bool } and return it.
	//   2. Check whether it implements the interface{ Unwrap() http.ResponseWriter } and retry 1.
	//   3. Return false instead.
	HTTPIsRespondedFunc = NewValueWithValidation(httpIsResponded, fhttprespR[bool]("HTTPIsResponded"))

	// GetHTTPStatusCodeFunc is used to get the status code of the http response.
	//
	// For the default implementation, it will try check http.ResponseWriter:
	//   1. Check whether it implements the interface{ StatusCode() int } and return it.
	//   2. Check whether it implements the interface{ Unwrap() http.ResponseWriter } and retry 1.
	//   3. Return 200 instead.
	GetHTTPStatusCodeFunc = NewValueWithValidation(getHTTPStatusCode, fhttprespR[int]("GetHTTPStatusCode"))

	// GetHTTPRequestFunc is used to get the http request from the request context.
	//
	// For the default implementation, it only supports the types or interfaces:
	//
	//	*http.Request
	//	interface{ Request() *http.Request }
	//	interface{ HTTPRequest() *http.Request }
	//	interface{ GetRequest() *http.Request }
	//	interface{ GetHTTPRequest() *http.Request }
	//
	// If not found, return nil.
	GetHTTPRequestFunc = NewValueWithValidation(getHTTPRequest, fActxAifaceR1[*http.Request]("GetHTTPRequest"))
)

// HTTPIsResponded is the proxy of HTTPIsRespondedFunc to call the function.
func HTTPIsResponded(ctx context.Context, w http.ResponseWriter, r *http.Request) bool {
	return HTTPIsRespondedFunc.Get()(ctx, w, r)
}

// GetHTTPStatusCode is the proxy of GetHTTPStatusCodeFunc to call the function.
func GetHTTPStatusCode(ctx context.Context, w http.ResponseWriter, r *http.Request) int {
	return GetHTTPStatusCodeFunc.Get()(ctx, w, r)
}

// GetHTTPRequest is the proxy of GetHTTPRequestFunc to call the function.
func GetHTTPRequest(ctx context.Context, req any) *http.Request {
	return GetHTTPRequestFunc.Get()(ctx, req)
}

func httpIsResponded(ctx context.Context, w http.ResponseWriter, r *http.Request) bool {
	for {
		switch _w := w.(type) {
		case interface{ WroteHeader() bool }:
			return _w.WroteHeader()
		case interface{ Unwrap() http.ResponseWriter }:
			w = _w.Unwrap()
		default:
			return false
		}
	}
}

func getHTTPStatusCode(ctx context.Context, w http.ResponseWriter, r *http.Request) int {
	for {
		switch _w := w.(type) {
		case interface{ StatusCode() int }:
			return _w.StatusCode()
		case interface{ Unwrap() http.ResponseWriter }:
			w = _w.Unwrap()
		default:
			return 200
		}
	}
}

func getHTTPRequest(ctx context.Context, req any) *http.Request {
	switch r := req.(type) {
	case *http.Request:
		return r

	case interface{ Request() *http.Request }:
		return r.Request()

	case interface{ HTTPRequest() *http.Request }:
		return r.HTTPRequest()

	case interface{ GetRequest() *http.Request }:
		return r.GetRequest()

	case interface{ GetHTTPRequest() *http.Request }:
		return r.GetHTTPRequest()

	default:
		return nil
	}
}
