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
	case int64, *int64:
		return val.(int64), true
	default:
		ival, err := strconv.ParseInt(
			fmt.Sprintf("%v", val),
			10, 64,
		)
		if err != nil {
			return 0, false
		}
		return ival, true
	}
}

// typecast from any to float
func AnyToFloat(val any) (float64, bool) {
	if val == nil {
		return 0.0, false
	}

	switch val.(type) {
	case float64, *float64:
		return val.(float64), true
	default:
		ival, err := strconv.ParseFloat(
			fmt.Sprintf("%v", val), 64,
		)
		if err != nil {
			return 0.0, false
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
	case *bool:
		if tp == nil {
			return false, false
		}
		return *tp, true
	default:
		ival, err := strconv.ParseBool(
			fmt.Sprintf("%v", val),
		)
		if err != nil {
			return false, false
		}
		return ival, true
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
	case reflect.Bool:
		if value == nil {
			return false, false
		}
		return AnyToBool(value)
	case reflect.Int, reflect.Int32, reflect.Int64:
		if value == nil {
			return 0, false
		}
		v, ok := AnyToInt(value)

		switch tp {
		case reflect.Int:
			return int(v), ok
		case reflect.Int32:
			return int32(v), ok
		default:
			return v, ok
		}
	case reflect.Float32, reflect.Float64:
		if value == nil {
			return 0, false
		}

		v, ok := AnyToFloat(value)
		if tp == reflect.Float32 {
			return float32(v), ok
		}

		return v, ok
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
