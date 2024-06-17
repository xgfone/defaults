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
)

var TrimPkgFile func(string) string

var (
	initfuncs []function
	exitfuncs []function
)

func OnInit(skip int, f func()) { initfuncs = register(initfuncs, "init", skip+1, f) }
func OnExit(skip int, f func()) { exitfuncs = register(exitfuncs, "init", skip+1, f) }

func RunInit() { runinits(initfuncs) }
func RunExit() { runexits(exitfuncs) }

type function struct {
	Func func()
	File string
	Line int
}

func (f function) runinit() { f.print("init"); f.Func() }
func (f function) runexit() { defer f.recover(); f.print("exit"); f.Func() }
func (f function) recover() {
	if r := recover(); r != nil {
		slog.Error("exit func panics", "file", f.File, "line", f.Line, "panic", r)
	}
}

func (f function) print(ftype string) {
	if DEBUG {
		slog.Info(fmt.Sprintf("run %s func", ftype), "file", f.File, "line", f.Line)
	}
}

func runinits(funcs []function) {
	for i := range funcs {
		funcs[i].runinit()
	}
}

func runexits(funcs []function) {
	for _len := len(funcs) - 1; _len >= 0; _len-- {
		funcs[_len].runexit()
	}
}

func register(funcs []function, ftype string, skip int, f func()) []function {
	if f == nil {
		panic(ftype + " function is nil")
	}

	file, line := getFileLine(skip + 2)
	funcs = append(funcs, function{Func: f, Line: line, File: file})
	return funcs
}

func getFileLine(skip int) (file string, line int) {
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		file = TrimPkgFile(file)
	} else {
		file = "???"
	}
	return
}
