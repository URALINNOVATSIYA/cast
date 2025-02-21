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
	if value == nil {
		return false, nil
	}
	switch v := value.(type) {
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
		return false, fmt.Errorf("failed to cast %q to bool", v)
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
			return AsType[bool](elemOf(rv.Elem()).Interface())
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
