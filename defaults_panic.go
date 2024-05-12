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
	"log/slog"
)

var (
	// HandlePanicFunc is used to handle the panic value returned by recover().
	HandlePanicFunc = NewValueWithValidation(handlePanic, fActxAiface("HandlePanic"))
)

// HandlePanic is the proxy of HandlePanicFunc to call the funciton.
func HandlePanic(ctx context.Context, r any) {
	HandlePanicFunc.Get()(ctx, r)
}

func handlePanic(ctx context.Context, r any) {
	slog.Error("wrap a panic", "panic", r, "stacks", GetStacks(2))
}

// Recover is a convenient function to wrap and recover the panic if occurring,
// then call HandlePanic to handle it.
//
// NOTICE: It must be called after defer, like
//
//	defer Recover(context.Background())
func Recover(ctx context.Context) {
	if r := recover(); r != nil {
		if ctx == nil {
			ctx = context.Background()
		}
		HandlePanic(ctx, r)
	}
}
