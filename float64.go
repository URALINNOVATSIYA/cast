package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToFloat64(value any) float64 {
	r, err := AsFloat64(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsFloat64(value any) (float64, error) {
	if value == nil {
		return 0, nil
	}
	switch v := value.(type) {
	case uint:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
			return AsType[float64](elemOf(rv.Elem()).Interface())
		}
	}
	return 0, fmt.Errorf("failed to cast %T to float64", value)
}

func asFloat64(value reflect.Value) (reflect.Value, error) {
	r, err := AsFloat64(value.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(r), nil
}
