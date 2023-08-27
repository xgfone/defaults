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
	// HeaderXRequestID is used by GetRequestIDFunc to try
	// to get the request id from the http request.
	HeaderXRequestID = "X-Request-Id"

	// GetRequestIDFunc is used to get the unique request session id.
	//
	// For the default implementation, it only supports the types and interfaces:
	//
	// 	*http.Request
	//	interface{ RequestID() string }
	//	interface{ GetRequestID() string }
	//
	// Or, return "".
	GetRequestIDFunc = NewValueWithValidation(getRequestID, fActxAifaceR1[string]("GetRequestID"))
)

// GetRequestID is the proxy of GetRequestIDFunc to call the function.
func GetRequestID(ctx context.Context, req interface{}) string {
	return GetRequestIDFunc.Get()(ctx, req)
}

func getRequestID(ctx context.Context, req interface{}) string {
	switch r := req.(type) {
	case interface{ RequestID() string }:
		return r.RequestID()

	case interface{ GetRequestID() string }:
		return r.GetRequestID()

	case *http.Request:
		return r.Header.Get(HeaderXRequestID)

	default:
		return ""
	}
}
