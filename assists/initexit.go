// Copyright 2024 xgfone
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
	"fmt"
	"log/slog"
	"runtime"
	"sync"
	"time"
)

var TrimPkgFile = func(file string) string { return file }

/// ----------------------------------------------------------------------- ///

var initfuncs []func()

// OnInit registers a function called when calling RunInit().
func OnInit(f func()) {
	initfuncs = append(initfuncs, f)
	_traceregister("init", 2)
}

// RunInit calls the init functions in turn.
func RunInit() {
	iter(initfuncs, func(f func()) { f() })
}

/// ----------------------------------------------------------------------- ///

var (
	exitfuncs  []func()
	cleanfuncs []func()
	exitonce   sync.Once
	exitedch   = make(chan struct{})
)

// OnExitPost registers a function called after calling exit functions.
func OnExitPost(f func()) {
	cleanfuncs = append(cleanfuncs, f)
	_traceregister("exitpost", 1)
}

// OnExit registers a function called when calling RunExit().
func OnExit(f func()) {
	exitfuncs = append(exitfuncs, f)
	_traceregister("exit", 2)
}

func WaitExit() { <-exitedch }

// RunExit calls the exit functions in reverse turn.
func RunExit() {
	exitonce.Do(exit)
	WaitExit()
}

func exit() {
	reverseIter(exitfuncs, runexit)
	reverseIter(cleanfuncs, runexit)
	close(exitedch)
}

func runexit(f func()) {
	defer exitrecover()
	f()
}

func exitrecover() {
	if r := recover(); r != nil {
		slog.Error("exit func panics", "panic", r)
	}
}

func init() { OnExitPost(func() { time.Sleep(time.Millisecond * 10) }) }

/// ----------------------------------------------------------------------- ///

func _traceregister(kind string, skip int) {
	if DEBUG {
		file, line := getfileline(skip + 2)
		msg := fmt.Sprintf("register %s function", kind)
		slog.Info(msg, "file", file, "line", line)
	}
}

func getfileline(skip int) (file string, line int) {
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		file = TrimPkgFile(file)
	} else {
		file = "???"
	}
	return
}
