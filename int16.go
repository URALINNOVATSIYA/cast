package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToInt16(value any) int16 {
	r, err := AsInt16(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsInt16(value any) (int16, error) {
	if value == nil {
		return 0, nil
	}
	switch v := value.(type) {
	case int:
		return int16(v), nil
	case uint:
		return int16(v), nil
	case int8:
		return int16(v), nil
	case uint8:
		return int16(v), nil
	case int16:
		return v, nil
	case uint16:
		return int16(v), nil
	case int32:
		return int16(v), nil
	case uint32:
		return int16(v), nil
	case int64:
		return int16(v), nil
	case uint64:
		return int16(v), nil
	case float32:
		return int16(v), nil
	case float64:
		return int16(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return int16(value), nil
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
			return AsType[int16](elemOf(rv.Elem()).Interface())
		}
	}
	return 0, fmt.Errorf("failed to cast %T to int16", value)
}

func asInt16(value reflect.Value) (reflect.Value, error) {
	r, err := AsInt16(value.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(r), nil
}
