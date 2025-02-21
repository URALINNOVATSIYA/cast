package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToUint32(value any) uint32 {
	r, err := AsUint32(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsUint32(value any) (uint32, error) {
	if value == nil {
		return 0, nil
	}
	switch v := value.(type) {
	case int:
		return uint32(v), nil
	case uint:
		return uint32(v), nil
	case int8:
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	case int16:
		return uint32(v), nil
	case uint16:
		return uint32(v), nil
	case int32:
		return uint32(v), nil
	case uint32:
		return v, nil
	case int64:
		return uint32(v), nil
	case uint64:
		return uint32(v), nil
	case float32:
		return uint32(v), nil
	case float64:
		return uint32(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		value, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return uint32(value), nil
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
			return AsType[uint32](elemOf(rv.Elem()).Interface())
		}
	}
	return 0, fmt.Errorf("failed to cast %T to uint32", value)
}

func asUint32(value reflect.Value) (reflect.Value, error) {
	r, err := AsUint32(value.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(r), nil
}
