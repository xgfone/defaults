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

import "log/slog"

var (
	// OnInitFunc is used to register the init function.
	OnInitFunc = NewValueWithValidation(oninit, fA1Validation[func()]("OnInit"))
)

// OnInit is the proxy of OnInitFunc to register the init function f.
func OnInit(f func()) { OnInitFunc.Get()(f) }

func oninit(f func()) {
	slog.Warn("system does not set the init function register", "caller", caller(1))
}
