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
	exitlock  sync.Mutex
	exitfuncs []func()
	exited    bool

	cleanfuncs []func()
)

// OnClean registers a function called after calling exit functions.
func OnClean(f func()) {
	cleanfuncs = append(cleanfuncs, f)
	_traceregister("clean", 1)
}

// OnExit registers a function called when calling RunExit().
func OnExit(f func()) {
	exitfuncs = append(exitfuncs, f)
	_traceregister("exit", 2)
}

// RunExit calls the exit functions in reverse turn.
func RunExit() {
	exitlock.Lock()
	defer exitlock.Unlock()
	if !exited {
		exited = true
		reverseIter(exitfuncs, runexit)
		reverseIter(cleanfuncs, runexit)
	}
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

func init() { OnClean(func() { time.Sleep(time.Millisecond * 10) }) }

/// ----------------------------------------------------------------------- ///

func _traceregister(kind string, skip int) {
	if DEBUG {
		file, line := getfileline(skip + 2)
		fmt.Printf("register %s function: file=%s, line=%d\n", kind, file, line)
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
