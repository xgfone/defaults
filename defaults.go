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
	"sync/atomic"
	"time"
)

// Pre-define some global variables about time.
var (
	TimeNow      = NewValueWithValidation(time.Now, validateTimeNow)
	TimeFormat   = NewValueWithValidation(time.RFC3339Nano, validateTimeFormat)
	TimeFormats  = NewValueWithValidation([]string{time.RFC3339Nano, "2006-01-02 15:04:05"}, validateTimeFormats)
	TimeLocation = NewValueWithValidation(time.UTC, validateTimeLocation)
)

func validateTimeNow(f func() time.Time) error {
	if f == nil {
		return errors.New("the time now function must not be nil")
	}
	return nil
}

func validateTimeFormat(s string) error {
	if s == "" {
		return errors.New("time format layout must not be empty")
	}
	return nil
}

func validateTimeLocation(loc *time.Location) error {
	if loc == nil {
		return errors.New("time location must not be nil")
	}
	return nil
}

func validateTimeFormats(ss []string) error {
	if len(ss) == 0 {
		return errors.New("time format layouts must not be empty")
	}
	for _, s := range ss {
		if s == "" {
			return errors.New("time format layouts must not be empty")
		}
	}
	return nil
}

// Now returns the current time by using TimeNow and TimeLocation.
func Now() time.Time { return TimeNow.Get()().In(TimeLocation.Get()) }

type valuer[T any] struct {
	Value T
}

// Value represents a common value.
type Value[T any] struct {
	verify func(T) error
	value  atomic.Value
}

// NewValue returns a new Value with the initial value.
func NewValue[T any](initial T) *Value[T] {
	return NewValueWithValidation(initial, nil)
}

// NewValueWithValidation returns a new Value with the initial value and the validation.
//
// validate may be nil, which is equal to always return nil.
func NewValueWithValidation[T any](initial T, validate func(T) error) *Value[T] {
	v := new(Value[T])
	v.verify = validate
	v.Set(initial)
	return v
}

// Get returns the inner value.
func (v *Value[T]) Get() T {
	return v.value.Load().(valuer[T]).Value
}

// Set sets the value to new thread-safely.
//
// It will panic if failing to validate the new value.
func (v *Value[T]) Set(new T) {
	if err := v.Validate(new); err != nil {
		panic(err)
	}
	v.value.Store(valuer[T]{Value: new})
}

// Validate validate whether the input value is valid.
func (v *Value[T]) Validate(value T) error {
	if v.verify == nil {
		return nil
	}
	return v.verify(value)
}
