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

package defaults

import "github.com/xgfone/go-toolkit/runtimex"

var (
	// IsZeroFunc is used to check whether a value is ZERO.
	//
	// For the default implementation, it also supports
	//    interface{ IsZero() bool }
	//
	IsZeroFunc = NewValueWithValidation(runtimex.IsZero, fA1R1Validation[any, bool]("IsZero"))
)

// IsZero is the proxy of IsZeroFunc to call it.
func IsZero(value any) bool { return IsZeroFunc.Get()(value) }
