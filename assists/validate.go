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

package assists

// StructValidator is used to validate whether a struct value is valid.
type StructValidator interface {
	Validate(structvalue any) error
}

var _ StructValidator = StructValidateFunc(nil)

// StructValidateFunc is a struct validation function.
type StructValidateFunc func(structvalue any) error

// Validate implements the interface Validator.
func (f StructValidateFunc) Validate(structvalue any) error {
	return f(structvalue)
}

// RuleValidator is used to validate whether a value conforms with the rule.
type RuleValidator interface {
	Validate(value any, rule string) error
}

var _ RuleValidator = RuleValidateFunc(nil)

// RuleValidateFunc is a rule validation function.
type RuleValidateFunc func(value any, rule string) error

// Validate implements the interface RuleValidator.
func (f RuleValidateFunc) Validate(value any, rule string) error {
	return f(value, rule)
}
