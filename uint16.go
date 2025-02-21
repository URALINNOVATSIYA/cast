package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToUint16(value any) uint16 {
	r, err := AsUint16(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsUint16(value any) (uint16, error) {
	if value == nil {
		return 0, nil
	}
	switch v := value.(type) {
	case int:
		return uint16(v), nil
	case uint:
		return uint16(v), nil
	case int8:
		return uint16(v), nil
	case uint8:
		return uint16(v), nil
	case int16:
		return uint16(v), nil
	case uint16:
		return v, nil
	case int32:
		return uint16(v), nil
	case uint32:
		return uint16(v), nil
	case int64:
		return uint16(v), nil
	case uint64:
		return uint16(v), nil
	case float32:
		return uint16(v), nil
	case float64:
		return uint16(v), nil
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
		return uint16(value), nil
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
			return AsType[uint16](elemOf(rv.Elem()).Interface())
		}
	}
	return 0, fmt.Errorf("failed to cast %T to uint16", value)
}

func asUint16(value reflect.Value) (reflect.Value, error) {
	r, err := AsUint16(value.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(r), nil
}
