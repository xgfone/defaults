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

// Value represents a common value.
type Value[T any] struct {
	verify func(T) error
	update func(T)
	value  T
}

// NewValue returns a new Value with the initial value.
func NewValue[T any](initial T) *Value[T] {
	return NewValueWithValidation(initial, nil)
}

// NewValueWithValidation returns a new Value with the initial value and the validation.
//
// validate may be nil, which is equal to always return nil.
func NewValueWithValidation[T any](initial T, validate func(T) error) *Value[T] {
	return &Value[T]{verify: validate, value: initial}
}

// Get returns the inner value.
func (v *Value[T]) Get() T { return v.value }

// Set sets the value to new.
//
// It will panic if failing to validate the new value.
func (v *Value[T]) Set(new T) {
	if err := v.Validate(new); err != nil {
		panic(err)
	}
	v.value = new
	logset()
	if v.update != nil {
		v.update(v.value)
	}
}

// Swap sets the value to new and returns the old value.
//
// It will panic if failing to validate the new value.
func (v *Value[T]) Swap(new T) (old T) {
	if err := v.Validate(new); err != nil {
		panic(err)
	}
	old = v.value
	v.value = new
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
