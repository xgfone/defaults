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

import "errors"

var (
	// HandlePanicFunc is used to handle the panic value returned by recover().
	//
	// Default: log.Printf("wrap a panic: %+v", recoverValue)
	HandlePanicFunc = NewValueWithValidation(handlePanic, handlePanicValidation)
)

// HandlePanic is the proxy of HandlePanicFunc to call the funciton.
func HandlePanic(r interface{}) {
	HandlePanicFunc.Get()(r)
}

func handlePanic(r interface{}) {
	logkv("wrap a panic", "panic", r, "stacks", GetStacks(2))
}

func handlePanicValidation(f func(interface{})) error {
	if f == nil {
		return errors.New("HandlePanic function must not be nil")
	}
	return nil
}
