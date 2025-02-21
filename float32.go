package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToFloat32(value any) float32 {
	r, err := AsFloat32(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsFloat32(value any) (float32, error) {
	if value == nil {
		return 0, nil
	}
	switch v := value.(type) {
	case uint:
		return float32(v), nil
	case int:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	case float32:
		return v, nil
	case float64:
		return float32(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		value, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return 0, err
		}
		return float32(value), nil
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
			return AsType[float32](elemOf(rv.Elem()).Interface())
		}
	}
	return 0, fmt.Errorf("failed to cast %T to float32", value)
}

func asFloat32(value reflect.Value) (reflect.Value, error) {
	r, err := AsFloat32(value.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(r), nil
}
