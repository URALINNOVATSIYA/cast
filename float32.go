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
	switch v := value.(type) {
	case nil:
		return 0, nil
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
		switch rv.Kind() {
		case reflect.Bool, reflect.String,
			reflect.Int, reflect.Uint,
			reflect.Int8, reflect.Uint8,
			reflect.Int16, reflect.Uint16,
			reflect.Int32, reflect.Uint32,
			reflect.Int64, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			return AsFloat32(rv.Convert(kind2types[rv.Kind()]).Interface())
		case reflect.Pointer, reflect.Interface:
			rv = rv.Elem()
			if !rv.IsValid() {
				return 0, nil
			}
			return AsFloat32(rv.Interface())
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
