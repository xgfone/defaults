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

	"github.com/xgfone/go-defaults/assists"
	"github.com/xgfone/go-toolkit/runtimex"
)

func loginfo(msg string, kvs ...any) {
	if assists.DEBUG {
		slog.Info(msg, kvs...)
	}
}

func logset() {
	loginfo("set the default", "caller", runtimex.Caller(0))
}

func logswap() {
	loginfo("swap the default", "caller", runtimex.Caller(0))
}
