package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToUint8(value any) uint8 {
	r, err := AsUint8(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsUint8(value any) (uint8, error) {
	switch v := value.(type) {
	case nil:
		return 0, nil
	case int:
		return uint8(v), nil
	case uint:
		return uint8(v), nil
	case int8:
		return uint8(v), nil
	case uint8:
		return v, nil
	case int16:
		return uint8(v), nil
	case uint16:
		return uint8(v), nil
	case int32:
		return uint8(v), nil
	case uint32:
		return uint8(v), nil
	case int64:
		return uint8(v), nil
	case uint64:
		return uint8(v), nil
	case float32:
		return uint8(v), nil
	case float64:
		return uint8(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		if value, err := strconv.ParseUint(v, 10, 64); err == nil {
			return uint8(value), nil
		}
	default:
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Bool, reflect.String,
			reflect.Int, reflect.Uint,
			reflect.Int8, reflect.Uint8,
			reflect.Int16, reflect.Uint16,
			reflect.Int32, reflect.Uint32,
			reflect.Int64, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			return AsUint8(rv.Convert(kind2types[rv.Kind()]).Interface())
		case reflect.Pointer, reflect.Interface:
			rv = rv.Elem()
			if !rv.IsValid() {
				return 0, nil
			}
			return AsUint8(rv.Interface())
		}
	}
	return 0, fmt.Errorf("failed to cast %T to uint8", value)
}

func asUint8(value reflect.Value) (reflect.Value, error) {
	r, err := AsUint8(value.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(r), nil
}
