package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToInt64(value any) int64 {
	r, err := AsInt64(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsInt64(value any) (int64, error) {
	switch v := value.(type) {
	case nil:
		return 0, nil
	case int:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint64:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		if value, err := strconv.ParseInt(v, 10, 64); err == nil {
			return value, nil
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
			return AsInt64(rv.Convert(kind2types[rv.Kind()]).Interface())
		case reflect.Pointer, reflect.Interface:
			rv = rv.Elem()
			if !rv.IsValid() {
				return 0, nil
			}
			return AsInt64(rv.Interface())
		}
	}
	return 0, fmt.Errorf("failed to cast %T to int64", value)
}

func asInt64(value reflect.Value) (reflect.Value, error) {
	r, err := AsInt64(value.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(r), nil
}
