// Copyright 2023~2024 xgfone
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
	"os"

	"github.com/xgfone/go-defaults/assists"
)

var (
	// ExitFunc is used to exit the program.
	//
	// Default: os.Exit
	ExitFunc = NewValueWithValidation(os.Exit, fA1Validation[int]("Exit"))

	// ExitContextFunc is used to get a context
	// which is cancelled before the program exits.
	//
	// Default: context.Background
	ExitContextFunc = NewValueWithValidation(context.Background, fR1Validation[context.Context]("ExitContext"))

	// ExitWaitFunc is used to wait until the program exit.
	//
	// Default: <-ExitContext().Done()
	ExitWaitFunc = NewValueWithValidation(exitwait, fValidation("ExitWait"))
)

// Exit is the proxy of ExitFunc to call the function to exit the program.
func Exit(code int) { ExitFunc.Get()(code) }

// ExitContext is the proxy of ExitContextFunc to call it to get a context
// which is cancelled before the program exits.
func ExitContext() context.Context { return ExitContextFunc.Get()() }

// ExitWait is the proxy of ExitContextFunc to call it
// to wait until the program exit.
func ExitWait() { ExitWaitFunc.Get()() }

func exitwait() { <-ExitContext().Done() }

// OnExit registers the exit function f, which is the proxy of assists.OnExit.
func OnExit(f func()) {
	assists.OnExit(1, f)
}
