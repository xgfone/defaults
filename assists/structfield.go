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

import (
	"reflect"
	"strings"
)

// StructFieldNameFuncWithTags returns a function to get the name and arg
// of struct field from the given tags in turn until a tag is found.
func StructFieldNameFuncWithTags(tags ...string) func(reflect.StructField) (name string, arg string) {
	return func(sf reflect.StructField) (name string, arg string) {
		for _, tag := range tags {
			if tag == "" {
				continue
			}

			value, ok := sf.Tag.Lookup(tag)
			if !ok {
				continue
			}

			if index := strings.IndexByte(value, ','); index > -1 {
				arg = strings.TrimSpace(value[index+1:])
				value = strings.TrimSpace(value[:index])
			}

			switch value {
			case "-":
			case "":
				name = sf.Name
			default:
				name = value
			}

			break
		}

		return
	}
}
