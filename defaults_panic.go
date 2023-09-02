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

import "context"

var (
	// HandlePanicContextFunc is used to handle the panic value returned by recover().
	HandlePanicContextFunc = NewValueWithValidation(handlePanicContext, fActxAiface("HandlePanicContext"))

	// HandlePanicFunc is used to handle the panic value returned by recover().
	//
	// Default: use HandlePanicContextFunc(context.Background(), r)
	//
	// DEPRECATED.
	HandlePanicFunc = NewValueWithValidation(handlePanic, fA1Validation[any]("HandlePanic"))
)

// HandlePanic is the proxy of HandlePanicFunc to call the funciton.
//
// DEPRECATED.
func HandlePanic(r any) {
	HandlePanicFunc.Get()(r)
}

// HandlePanicContext is the proxy of HandlePanicContextFunc to call the funciton.
func HandlePanicContext(c context.Context, r any) {
	HandlePanicContextFunc.Get()(c, r)
}

func handlePanic(r any) {
	HandlePanicContextFunc.Get()(context.Background(), r)
}

func handlePanicContext(c context.Context, r any) {
	logkv("wrap a panic", "panic", r, "stacks", GetStacks(2))
}
