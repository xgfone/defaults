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
	"reflect"

	"github.com/xgfone/go-defaults/assists"
)

var (
	// StructFieldNameFunc is used to get the obtain the name and arg
	// of the struct field, which needs to return a empty string for
	// the field name, that's "", if the field should be ignored.
	//
	// Example:
	//
	//	StructFieldNameFunc.Set(func(sf reflect.StructField) (name string, arg string) {
	//	    value := sf.Tag.Get("json")
	//	    if index := strings.IndexByte(value, ','); index > -1 {
	//	        arg = strings.TrimSpace(value[index+1:])
	//	        value = strings.TrimSpace(value[:index])
	//	    }
	//
	//	    switch value {
	//	    case "-":
	//	    case "":
	//	        name = sf.Name
	//	    default:
	//	        name = value
	//	    }
	//
	//	    return
	//	})
	//
	StructFieldNameFunc = NewValueWithValidation(assists.StructFieldNameFuncWithTags("json"),
		fA1R2Validation[reflect.StructField, string, string]("StructFieldName"))
)

// GetStructFieldName is the proxy of StructFieldNameFunc to call the function,
// just like StructFieldNameFunc.Get()(sf).
func GetStructFieldName(sf reflect.StructField) (name, arg string) {
	return StructFieldNameFunc.Get()(sf)
}
