// Copyright 2023~2024 xgfone
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
	"time"

	"github.com/xgfone/go-toolkit/timex"
)

var (
	// DEPRECATED!!! Please use timex.Format instead.
	TimeFormat = NewValue(timex.Format)

	// DEPRECATED!!! Please use timex.Formats instead.
	TimeFormats = NewValue(timex.Formats)

	// DEPRECATED!!! Please use timex.Now instead.
	TimeNowFunc = NewValue(timex.Now)

	// DEPRECATED!!! Please use timex.Location instead.
	TimeLocation = NewValue(timex.Location)
)

func init() {
	TimeFormat.update = func(new string) { timex.Format = new }
	TimeFormats.update = func(new []string) { timex.Formats = new }
	TimeNowFunc.update = func(new func() time.Time) { timex.Now = new }
	TimeLocation.update = func(new *time.Location) { timex.Location = new }
}

// Now is eqaul to timex.Now.
//
// DEPRECATED!!! Please use timex.Now instead.
func Now() time.Time { return timex.Now() }

// Unix is eqaul to timex.Unix.
//
// DEPRECATED!!! Please use timex.Unix instead.
func Unix(sec, nsec int64) time.Time { return timex.Unix(sec, nsec) }

// Today is eqaul to timex.Today.
//
// DEPRECATED!!! Please use timex.Today instead.
func Today() time.Time { return timex.Today() }
