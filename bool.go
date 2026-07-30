package cast

import (
	"fmt"
	"reflect"
)

func ToBool(value any) bool {
	r, err := AsBool(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsBool(value any) (bool, error) {
	switch v := value.(type) {
	case nil:
		return false, nil
	case bool:
		return v, nil
	case int:
		return v != 0, nil
	case uint:
		return v != 0, nil
	case int8:
		return v != 0, nil
	case uint8:
		return v != 0, nil
	case int16:
		return v != 0, nil
	case uint16:
		return v != 0, nil
	case int32:
		return v != 0, nil
	case uint32:
		return v != 0, nil
	case int64:
		return v != 0, nil
	case uint64:
		return v != 0, nil
	case float32:
		return v != 0, nil
	case float64:
		return v != 0, nil
	case string:
		if v == "" || v == "false" {
			return false, nil
		}
		if v == "true" {
			return true, nil
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
			return AsBool(rv.Convert(kind2types[rv.Kind()]).Interface())
		case reflect.Pointer, reflect.Interface:
			rv = rv.Elem()
			if !rv.IsValid() {
				return false, nil
			}
			return AsBool(rv.Interface())
		}
	}
	return false, fmt.Errorf("failed to cast %T to bool", value)
}

func asBool(value reflect.Value) (reflect.Value, error) {
	r, err := AsBool(value.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(r), nil
}
