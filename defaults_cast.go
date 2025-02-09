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
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/xgfone/go-toolkit/timex"
)

var (
	// ToBoolFunc is used to convert an input to bool.
	ToBoolFunc = NewValueWithValidation(tobool, castValidation[bool]("ToBool"))

	// ToInt64Func is used to convert an input to int64.
	ToInt64Func = NewValueWithValidation(toint64, castValidation[int64]("ToInt64"))

	// ToUint64Func is used to convert an input to uint64.
	ToUint64Func = NewValueWithValidation(touint64, castValidation[uint64]("ToUint64"))

	// ToFloat64Func is used to convert an input to float64.
	ToFloat64Func = NewValueWithValidation(tofloat64, castValidation[float64]("ToFloat64"))

	// ToStringFunc is used to convert an input to string.
	ToStringFunc = NewValueWithValidation(tostring, castValidation[string]("ToString"))

	// ToDurationFunc is used to convert an input to time.Duraiton.
	ToDurationFunc = NewValueWithValidation(toduration, castValidation[time.Duration]("ToDuration"))

	// ToTimeFunc is used to convert an input to time.Time.
	ToTimeFunc = NewValueWithValidation(totime, castValidation[time.Time]("ToTime"))
)

// ToBool is the proxy of ToBoolFunc to convert an input to bool.
func ToBool(input any) (bool, error) { return ToBoolFunc.Get()(input) }

// ToInt64 is the proxy of ToInt64Func to convert an input to int64.
func ToInt64(input any) (int64, error) { return ToInt64Func.Get()(input) }

// ToUint64 is the proxy of ToUint64Func to convert an input to uint64.
func ToUint64(input any) (uint64, error) { return ToUint64Func.Get()(input) }

// ToFloat64 is the proxy of ToFloat64Func to convert an input to float64.
func ToFloat64(input any) (float64, error) { return ToFloat64Func.Get()(input) }

// ToString is the proxy of ToStringFunc to convert an input to string.
func ToString(input any) (string, error) { return ToStringFunc.Get()(input) }

// ToDuration is the proxy of ToDurationFunc to convert an input to time.Duration.
func ToDuration(input any) (time.Duration, error) { return ToDurationFunc.Get()(input) }

// ToTime is the proxy of ToTimeFunc to convert an input to time.Time.
func ToTime(input any) (time.Time, error) { return ToTimeFunc.Get()(input) }

func tobool(src any) (dst bool, err error) {
	switch src := src.(type) {
	case nil:
	case bool:
		dst = src
	case string:
		if src != "" {
			dst, err = strconv.ParseBool(src)
		}
	case []byte:
		switch len(src) {
		case 0:
		case 1:
			switch src[0] {
			case '\x00':
			case '\x01':
				dst = true
			default:
				err = fmt.Errorf("unsupport to convert []byte to bool")
			}
		default:
			err = fmt.Errorf("unsupport to convert []byte to bool")
		}
	case float32:
		dst = src != 0
	case float64:
		dst = src != 0
	case int:
		dst = src != 0
	case int8:
		dst = src != 0
	case int16:
		dst = src != 0
	case int32:
		dst = src != 0
	case int64:
		dst = src != 0
	case uint:
		dst = src != 0
	case uint8:
		dst = src != 0
	case uint16:
		dst = src != 0
	case uint32:
		dst = src != 0
	case uint64:
		dst = src != 0
	case uintptr:
		dst = src != 0
	case interface{ Bool() bool }:
		dst = src.Bool()
	case interface{ IsZero() bool }:
		dst = !src.IsZero()
	default:
		err = fmt.Errorf("unsupport to convert %#T to bool", src)
	}
	return
}

func toint64(src any) (dst int64, err error) {
	switch src := src.(type) {
	case nil:
	case bool:
		if src {
			dst = 1
		}
	case string:
		if src != "" {
			dst, err = strconv.ParseInt(src, 0, 64)
		}
	case []byte:
		if len(src) > 0 {
			dst, err = strconv.ParseInt(string(src), 0, 64)
		}
	case float32:
		dst = int64(src)
	case float64:
		dst = int64(src)
	case int:
		dst = int64(src)
	case int8:
		dst = int64(src)
	case int16:
		dst = int64(src)
	case int32:
		dst = int64(src)
	case int64:
		dst = src
	case uint:
		dst = int64(src)
	case uint8:
		dst = int64(src)
	case uint16:
		dst = int64(src)
	case uint32:
		dst = int64(src)
	case uint64:
		dst = int64(src)
	case uintptr:
		dst = int64(src)
	case time.Duration:
		dst = int64(src / time.Millisecond)
	case *time.Duration:
		dst = int64(*src / time.Millisecond)
	case time.Time:
		dst = src.Unix()
	case *time.Time:
		dst = src.Unix()
	case interface{ Int64() int64 }:
		dst = src.Int64()
	case interface{ Int() int64 }:
		dst = src.Int()
	default:
		err = fmt.Errorf("unsupport to convert %#T to int64", src)
	}
	return
}

func touint64(src any) (dst uint64, err error) {
	switch src := src.(type) {
	case nil:
	case bool:
		if src {
			dst = 1
		}
	case string:
		if src != "" {
			dst, err = strconv.ParseUint(src, 0, 64)
		}
	case []byte:
		if len(src) > 0 {
			dst, err = strconv.ParseUint(string(src), 0, 64)
		}
	case float32:
		if src < 0 {
			err = errors.New("cannot convert a negative to uint64")
		} else {
			dst = uint64(src)
		}
	case float64:
		if src < 0 {
			err = errors.New("cannot convert a negative to uint64")
		} else {
			dst = uint64(src)
		}
	case int:
		if src < 0 {
			err = errors.New("cannot convert a negative to uint64")
		} else {
			dst = uint64(src)
		}
	case int8:
		if src < 0 {
			err = errors.New("cannot convert a negative to uint64")
		} else {
			dst = uint64(src)
		}
	case int16:
		if src < 0 {
			err = errors.New("cannot convert a negative to uint64")
		} else {
			dst = uint64(src)
		}
	case int32:
		if src < 0 {
			err = errors.New("cannot convert a negative to uint64")
		} else {
			dst = uint64(src)
		}
	case int64:
		if src < 0 {
			err = errors.New("cannot convert a negative to uint64")
		} else {
			dst = uint64(src)
		}
	case uint:
		dst = uint64(src)
	case uint8:
		dst = uint64(src)
	case uint16:
		dst = uint64(src)
	case uint32:
		dst = uint64(src)
	case uint64:
		dst = src
	case uintptr:
		dst = uint64(src)
	case interface{ Uint64() uint64 }:
		dst = src.Uint64()
	case interface{ Uint() uint64 }:
		dst = src.Uint()
	default:
		err = fmt.Errorf("unsupport to convert %#T to uint64", src)
	}
	return
}

func tofloat64(src any) (dst float64, err error) {
	switch src := src.(type) {
	case nil:
	case bool:
		if src {
			dst = 1
		}
	case string:
		if src != "" {
			dst, err = strconv.ParseFloat(src, 64)
		}
	case []byte:
		if len(src) > 0 {
			dst, err = strconv.ParseFloat(string(src), 64)
		}
	case float32:
		dst = float64(src)
	case float64:
		dst = src
	case int:
		dst = float64(src)
	case int8:
		dst = float64(src)
	case int16:
		dst = float64(src)
	case int32:
		dst = float64(src)
	case int64:
		dst = float64(src)
	case uint:
		dst = float64(src)
	case uint8:
		dst = float64(src)
	case uint16:
		dst = float64(src)
	case uint32:
		dst = float64(src)
	case uint64:
		dst = float64(src)
	case uintptr:
		dst = float64(src)
	case time.Duration:
		dst = float64(src) / float64(time.Second)
	case *time.Duration:
		dst = float64(*src) / float64(time.Second)
	case interface{ Float64() float64 }:
		dst = src.Float64()
	case interface{ Float() float64 }:
		dst = src.Float()
	default:
		err = fmt.Errorf("unsupport to convert %#T to float64", src)
	}
	return
}

func tostring(src any) (dst string, err error) {
	switch src := src.(type) {
	case nil:
	case bool:
		dst = strconv.FormatBool(src)
	case string:
		dst = src
	case []byte:
		dst = string(src)
	case float32:
		dst = strconv.FormatFloat(float64(src), 'f', -1, 32)
	case float64:
		dst = strconv.FormatFloat(src, 'f', -1, 64)
	case int:
		dst = strconv.FormatInt(int64(src), 10)
	case int8:
		dst = strconv.FormatInt(int64(src), 10)
	case int16:
		dst = strconv.FormatInt(int64(src), 10)
	case int32:
		dst = strconv.FormatInt(int64(src), 10)
	case int64:
		dst = strconv.FormatInt(src, 10)
	case uint:
		return strconv.FormatUint(uint64(src), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(src), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(src), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(src), 10), nil
	case uint64:
		return strconv.FormatUint(src, 10), nil
	case uintptr:
		return strconv.FormatUint(uint64(src), 10), nil
	case time.Time:
		dst = src.Format(time.RFC3339Nano)
	case *time.Time:
		dst = src.Format(time.RFC3339Nano)
	case error:
		dst = src.Error()
	case fmt.Stringer:
		dst = src.String()
	default:
		err = fmt.Errorf("unsupport to convert %#T to string", src)
	}
	return
}

func toduration(src any) (dst time.Duration, err error) {
	switch src := src.(type) {
	case nil:
	case string:
		dst, err = parseDuration(src)
	case []byte:
		dst, err = parseDuration(string(src))
	case float32:
		dst = time.Duration(float64(src) * float64(time.Second))
	case float64:
		dst = time.Duration(src * float64(time.Second))
	case int:
		dst = time.Duration(src) * time.Millisecond
	case int8:
		dst = time.Duration(src) * time.Millisecond
	case int16:
		dst = time.Duration(src) * time.Millisecond
	case int32:
		dst = time.Duration(src) * time.Millisecond
	case int64:
		dst = time.Duration(src) * time.Millisecond
	case uint:
		dst = time.Duration(src) * time.Millisecond
	case uint8:
		dst = time.Duration(src) * time.Millisecond
	case uint16:
		dst = time.Duration(src) * time.Millisecond
	case uint32:
		dst = time.Duration(src) * time.Millisecond
	case uint64:
		dst = time.Duration(src) * time.Millisecond
	case uintptr:
		dst = time.Duration(src) * time.Millisecond
	case time.Duration:
		dst = src
	case *time.Duration:
		dst = *src
	case interface{ Duration() time.Duration }:
		dst = src.Duration()
	default:
		err = fmt.Errorf("unsupport to convert %#T to time.Duration", src)
	}
	return
}

func totime(src any) (dst time.Time, err error) {
	loc := timex.Location
	switch src := src.(type) {
	case nil:
		dst = dst.In(loc)
	case string:
		dst, err = parseTime(src)
	case []byte:
		dst, err = parseTime(string(src))
	case float32:
		dst = time.Unix(int64(src), 0).In(loc)
	case float64:
		dst = time.Unix(int64(src), 0).In(loc)
	case int:
		dst = time.Unix(int64(src), 0).In(loc)
	case int32:
		dst = time.Unix(int64(src), 0).In(loc)
	case int64:
		dst = time.Unix(int64(src), 0).In(loc)
	case uint:
		dst = time.Unix(int64(src), 0).In(loc)
	case uint32:
		dst = time.Unix(int64(src), 0).In(loc)
	case uint64:
		dst = time.Unix(int64(src), 0).In(loc)
	case time.Time:
		dst = src.In(loc)
	case *time.Time:
		dst = src.In(loc)
	case interface{ Time() time.Time }:
		dst = src.Time()
	default:
		err = fmt.Errorf("unsupport to convert %#T to time.Time", src)
	}
	return
}

func parseDuration(src string) (dst time.Duration, err error) {
	_len := len(src)
	if _len == 0 {
		return
	}

	switch src[_len-1] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		var i int64
		i, err = strconv.ParseInt(src, 10, 64)
		dst = time.Duration(i) * time.Millisecond
	default:
		dst, err = time.ParseDuration(src)
	}

	return
}

func parseTime(value string) (time.Time, error) {
	loc := timex.Location
	switch value {
	case "", "0000-00-00 00:00:00", "0000-00-00 00:00:00.000", "0000-00-00 00:00:00.000000":
		return time.Time{}.In(loc), nil
	}

	if isIntegerString(value) {
		i, err := strconv.ParseInt(value, 10, 64)
		return Unix(i, 0), err
	}

	for _, layout := range timex.Formats {
		if t, err := time.ParseInLocation(layout, value, loc); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time '%s'", value)
}

func isIntegerString(s string) bool {
	_len := len(s)
	if _len == 0 {
		return false
	}

	switch s[0] {
	case '-', '+':
		s = s[1:]
		_len--
	}

	for i := 0; i < _len; i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			return false
		}
	}
	return true
}

func castValidation[T any](name string) func(func(any) (T, error)) error {
	return fA1R2Validation[any, T, error](name)
}
