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
	"errors"
	"fmt"
	"runtime"
	"strings"
)

var (
	// TrimPkgFileFunc is used to trim the prefix of the package file path.
	TrimPkgFileFunc = NewValueWithValidation(trimPkgFile, trimPkgFileValidation)

	// GetStacksFunc is used to get the stacks of the function calling.
	GetStacksFunc = NewValueWithValidation(getStacks, getStacksValidation)
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

		frame.File = TrimPkgFile(frame.File)
		if frame.Function == "" {
			stacks = append(stacks, fmt.Sprintf("%s:%d", frame.File, frame.Line))
		} else {
			name := frame.Function
			if index := strings.LastIndexByte(frame.Function, '.'); index > -1 {
				name = frame.Function[index+1:]
			}
			stacks = append(stacks, fmt.Sprintf("%s:%s:%d", frame.File, name, frame.Line))
		}
	}

	return stacks
}

func getStacksValidation(f func(skip int) []string) error {
	if f == nil {
		return errors.New("GetStacks function must not be nil")
	}
	return nil
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

func trimPkgFileValidation(f func(file string) string) error {
	if f == nil {
		return errors.New("TrimPkgFile function must not be nil")
	}
	return nil
}
