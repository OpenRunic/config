package reader

import (
	"fmt"
	"maps"
	"reflect"
	"strconv"
	"strings"
)

// typecast from any to int
func AnyToInt(val any) (int64, bool) {
	if val == nil {
		return 0, false
	}

	switch val.(type) {
	case int, int32:
		return val.(int64), true
	default:
		v := fmt.Sprintf("%v", val)
		ival, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, false
		}
		return ival, true
	}
}

// typecast from any to bool
func AnyToBool(val any) (bool, bool) {
	if val == nil {
		return false, false
	}

	switch tp := val.(type) {
	case bool:
		return tp, true
	case string:
		return tp == "true", true
	case int:
		return tp != 0, true
	case *bool:
		if tp == nil {
			return false, false
		}
		return *tp, true
	case *string:
		if tp == nil {
			return false, false
		}
		return *tp == "true", true
	case *int:
		if tp == nil {
			return false, false
		}
		return *tp != 0, true
	default:
		return false, false
	}
}

// cast value to type if needed
func CastValue(value any, tp reflect.Kind) (any, bool) {
	if value != nil {
		kind := reflect.ValueOf(value).Kind()
		if kind == tp {
			return value, true
		}
	}

	switch tp {
	case reflect.String:
		if value == nil {
			return "", false
		}
		return fmt.Sprintf("%v", value), true
	case reflect.Int:
		if value == nil {
			return 0, false
		}
		return AnyToInt(value)
	case reflect.Bool:
		if value == nil {
			return false, false
		}
		return AnyToBool(value)
	}

	return nil, false
}

// cast value to slice of type
func CastSliceValue(value any, tp reflect.Kind) ([]any, bool) {
	if value != nil {
		vtp := reflect.ValueOf(value)
		if vtp.Kind() == reflect.Slice {
			if vtp.Elem().Kind() == tp {
				return value.([]any), true
			}

			return nil, false
		}
	}

	if value == nil {
		return nil, false
	}

	data := make([]any, 0)

	switch tp {
	case reflect.String:
		splits := strings.Split(fmt.Sprintf("%v", value), ",")
		for _, s := range splits {
			data = append(data, s)
		}

		return data, true
	}

	return nil, false
}

// flatten nest map data to dot separated key data
func FlattenMap(data map[string]any, prefix string) map[string]any {
	result := make(map[string]any, 0)

	for key, val := range data {
		if reflect.Map == reflect.ValueOf(val).Kind() {
			tResult := FlattenMap(val.(map[string]any), key+".")
			maps.Copy(result, tResult)
		} else {
			result[prefix+key] = val
		}
	}

	return result
}

// convert string to snake case
func ToSnakeCase(s string) string {
	lastSep := false
	var result string

	for i, v := range s {
		if v < 65 {
			if !lastSep {
				result += "_"
				lastSep = true
			}
		} else {
			if i > 0 && v >= 'A' && v <= 'Z' {
				result += "_"
				lastSep = true
			} else {
				lastSep = false
			}

			result += string(v)
		}
	}

	return strings.ToLower(result)
}
