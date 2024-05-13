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

//go:build atomic

package defaults

import "sync/atomic"

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
	logset()
}

// Swap sets the value to new thread-safely and returns the old value.
//
// It will panic if failing to validate the new value.
func (v *Value[T]) Swap(new T) (old T) {
	if err := v.Validate(new); err != nil {
		panic(err)
	}
	if value := v.value.Swap(valuer[T]{Value: new}); value != nil {
		old = value.(valuer[T]).Value
	}
	logswap()
	return
}

// Validate validate whether the input value is valid.
func (v *Value[T]) Validate(value T) error {
	if v.verify == nil {
		return nil
	}
	return v.verify(value)
}
