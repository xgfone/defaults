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

import (
	"log/slog"
	"os"
	"strconv"
)

var isdebug bool

func init() {
	isdebug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
}

func logwarn(msg string, kvs ...any) {
	slog.Warn(msg, kvs...)
}

func logset() {
	if isdebug {
		slog.Info("set the default", "caller", GetCaller(2))
	}
}

func logswap() {
	if isdebug {
		slog.Info("swap the default", "caller", GetCaller(2))
	}
}
