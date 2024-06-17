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
	"fmt"
	"runtime"
	"strings"
)

var (
	// TrimPkgFileFunc is used to trim the prefix of the package file path.
	TrimPkgFileFunc = NewValueWithValidation(trimPkgFile, fA1R1Validation[string, string]("TrimPkgFile"))

	// GetStacksFunc is used to get the stacks of the function calling.
	GetStacksFunc = NewValueWithValidation(getStacks, fA1R1Validation[int, []string]("GetStacks"))

	GetCallerFunc = NewValueWithValidation(getCaller, fA1R1Validation[int, string]("GetCaller"))
)

// GetStacks is the proxy of GetStacksFunc to call the funciton.
func GetStacks(skip int) []string {
	return GetStacksFunc.Get()(skip + 3)
}

func getStacks(skip int) []string {
	var pcs [64]uintptr
	n := runtime.Callers(skip, pcs[:])
	if n == 0 {
		return nil
	}

	stacks := make([]string, 0, n)
	frames := runtime.CallersFrames(pcs[:n])
	for {
		frame, more := frames.Next()
		if !more {
			break
		}
		stacks = append(stacks, fmtframe(frame))
	}

	return stacks
}

// TrimPkgFile is the proxy of TrimPkgFileFunc to call the funciton.
func TrimPkgFile(file string) string {
	return TrimPkgFileFunc.Get()(file)
}

var trimPrefixes = []string{"/pkg/mod/", "/src/"}

func trimPkgFile(file string) string {
	for _, mark := range trimPrefixes {
		if index := strings.Index(file, mark); index > -1 {
			file = file[index+len(mark):]
			break
		}
	}
	return file
}

func GetCaller(skip int) string {
	return GetCallerFunc.Get()(skip + 3)
}

func getCaller(skip int) string {
	pcs := make([]uintptr, 1)
	if n := runtime.Callers(skip, pcs); n > 0 {
		frame, _ := runtime.CallersFrames(pcs).Next()
		if frame.PC != 0 {
			return fmtframe(frame)
		}
	}

	return "???"
}

func fmtframe(frame runtime.Frame) string {
	frame.File = TrimPkgFile(frame.File)
	if frame.Function == "" {
		return fmt.Sprintf("%s:%d", frame.File, frame.Line)
	}

	name := frame.Function
	if index := strings.LastIndexByte(frame.Function, '.'); index > -1 {
		name = frame.Function[index+1:]
	}

	return fmt.Sprintf("%s:%s:%d", frame.File, name, frame.Line)
}
