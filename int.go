package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToInt(value any) int {
	r, err := AsInt(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsInt(value any) (int, error) {
	switch v := value.(type) {
	case nil:
		return 0, nil
	case int:
		return v, nil
	case uint:
		return int(v), nil
	case int8:
		return int(v), nil
	case uint8:
		return int(v), nil
	case int16:
		return int(v), nil
	case uint16:
		return int(v), nil
	case int32:
		return int(v), nil
	case uint32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint64:
		return int(v), nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		if r, err := strconv.Atoi(v); err == nil {
			return r, nil
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
			return AsInt(rv.Convert(kind2types[rv.Kind()]).Interface())
		case reflect.Pointer, reflect.Interface:
			rv = rv.Elem()
			if !rv.IsValid() {
				return 0, nil
			}
			return AsInt(rv.Interface())
		}
	}
	return 0, fmt.Errorf("failed to cast %T to int", value)
}

func asInt(value reflect.Value) (reflect.Value, error) {
	r, err := AsInt(value.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(r), nil
}
