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
	// Default: calling assists.RunExit and os.Exit in turn.
	ExitFunc = NewValueWithValidation(exit, fA1Validation[int]("Exit"))

	// ExitWaitFunc is used to wait until the program exit.
	//
	// Default: assists.WaitExit
	ExitWaitFunc = NewValueWithValidation(assists.WaitExit, fValidation("ExitWait"))

	// ExitContextFunc is used to get the exit context.
	//
	// Default: assists.ExitContext
	ExitContextFunc = NewValueWithValidation(assists.ExitContext, fR1Validation[context.Context]("ExitContext"))

	// ExitSignalsFunc is used to get the signals to let the program exit.
	//
	// For default, on Unix/Linux or Windows, it contains the signals as follow:
	//
	//	os.Interrupt
	//	syscall.SIGTERM
	//	syscall.SIGQUIT
	//	syscall.SIGABRT
	//	syscall.SIGINT
	//
	// On others, it only contains the signal os.Interrupt.
	ExitSignalsFunc = NewValueWithValidation(exitSignalsFunc, fR1Validation[[]os.Signal]("ExitSignals"))
)

func exit(code int) {
	assists.RunExit()
	os.Exit(code)
}

// Exit is the proxy of ExitFunc to call the function to exit the program.
func Exit(code int) { ExitFunc.Get()(code) }

// ExitContext is the proxy of ExitContextFunc to get the exit context.
func ExitContext() context.Context { return ExitContextFunc.Get()() }

// ExitSignals is the proxy of ExitSignalsFunc to call it to get the exit signals.
func ExitSignals() []os.Signal { return ExitSignalsFunc.Get()() }

func exitSignalsFunc() []os.Signal { return exitsignals }

// ExitWait is the proxy of ExitContextFunc to call it
// to wait until the program exit.
func ExitWait() { ExitWaitFunc.Get()() }

// OnExit registers the exit function f, which is the proxy of assists.OnExit.
func OnExit(f func()) { assists.OnExit(f) }

// OnExitPost registers the post-exit function f, which is the proxy of assists.OnExitPost.
func OnExitPost(f func()) { assists.OnExitPost(f) }
