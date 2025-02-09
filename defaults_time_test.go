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
	"fmt"
	"testing"
	"time"

	"github.com/xgfone/go-toolkit/timex"
)

func TestLocation(t *testing.T) {
	timex.Location = time.UTC
	TimeLocation.Set(time.Local)

	if timex.Location != time.Local {
		t.Errorf("expect Location %s, but got %s", time.Local.String(), timex.Location.String())
	}
}

func TestToday(t *testing.T) {
	today := Today()
	expect := fmt.Sprintf("%04d-%02d-%02d 00:00:00",
		today.Year(), today.Month(), today.Day())

	if s := Today().Format(time.DateTime); s != expect {
		t.Errorf("expect '%s', but got '%s'", expect, s)
	}
}
