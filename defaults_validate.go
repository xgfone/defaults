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

import (
	"github.com/xgfone/defaults/assists"
)

var (
	// RuleValidator is used to validate whether a value conforms with the rule.
	//
	// Note: in general, it is used to validate a struct field with the tag rule.
	RuleValidator = NewValue(assists.RuleValidator(nil))

	// StructValidator is used to validate whether a struct value is valid.
	StructValidator = NewValue(assists.StructValidator(nil))
)

// ValidateStruct uses Validator to validate the struct value
// if StructValidator is not nil.
func ValidateStruct(value interface{}) (err error) {
	if v := StructValidator.Get(); v != nil {
		err = v.Validate(value)
	}
	return
}

// ValidateWithRule uses RuleValidator to validate a value conforms
// with the rule if RuleValidator is not nil.
func ValidateWithRule(value interface{}, rule string) (err error) {
	if v := RuleValidator.Get(); v != nil {
		err = v.Validate(value, rule)
	}
	return
}
