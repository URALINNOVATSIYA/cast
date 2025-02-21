package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToUint(value any) uint {
	r, err := AsUint(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsUint(value any) (uint, error) {
	if value == nil {
		return 0, nil
	}
	switch v := value.(type) {
	case int:
		return uint(v), nil
	case uint:
		return v, nil
	case int8:
		return uint(v), nil
	case uint8:
		return uint(v), nil
	case int16:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case int32:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case int64:
		return uint(v), nil
	case uint64:
		return uint(v), nil
	case float32:
		return uint(v), nil
	case float64:
		return uint(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		value, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return uint(value), nil
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
			return AsType[uint](elemOf(rv.Elem()).Interface())
		}
	}
	return 0, fmt.Errorf("failed to cast %T to uint", value)
}

func asUint(value reflect.Value) (reflect.Value, error) {
	r, err := AsUint(value.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(r), nil
}
