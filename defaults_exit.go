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

import "os"

var (
	// ExitFunc is used to exit the program.
	//
	// Default: os.Exit
	ExitFunc = NewValueWithValidation(os.Exit, fA1Validation[int]("Exit"))
)

// Exit is the proxy of ExitFunc to call the function to exit the program.
func Exit(code int) { ExitFunc.Get()(code) }
