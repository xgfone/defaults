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

// Package defaults provides some global default values.
package defaults

import (
	"errors"
	"time"
)

// TimeNow is the alias of TimeNowFunc.
//
// DEPRECATED! Please use TimeNowFunc instead.
var TimeNow = TimeNowFunc

// Pre-define some global variables about time.
var (
	TimeFormat   = NewValueWithValidation(time.RFC3339Nano, validateTimeFormat)
	TimeFormats  = NewValueWithValidation([]string{time.RFC3339Nano, "2006-01-02 15:04:05"}, validateTimeFormats)
	TimeLocation = NewValueWithValidation(time.UTC, validateTimeLocation)
	TimeNowFunc  = NewValueWithValidation(time.Now, validateTimeNow)
)

// Now returns the current time by using TimeNow and TimeLocation.
func Now() time.Time { return TimeNowFunc.Get()().In(TimeLocation.Get()) }

func validateTimeNow(f func() time.Time) error {
	if f == nil {
		return errors.New("TimeNow: the time now function must not be nil")
	}
	return nil
}

func validateTimeFormat(s string) error {
	if s == "" {
		return errors.New("TimeFormat: time format layout must not be empty")
	}
	return nil
}

func validateTimeLocation(loc *time.Location) error {
	if loc == nil {
		return errors.New("TimeLocation: time location must not be nil")
	}
	return nil
}

func validateTimeFormats(ss []string) error {
	if len(ss) == 0 {
		return errors.New("TimeFormats: time format layouts must not be empty")
	}
	for _, s := range ss {
		if s == "" {
			return errors.New("TimeFormats: time format layouts must not be empty")
		}
	}
	return nil
}
