package cast

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToString(value any) string {
	r, err := AsString(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsString(value any) (string, error) {
	switch v := value.(type) {
	case nil:
		return "", nil
	case string:
		return v, nil
	case bool:
		if v {
			return "true", nil
		}
		return "false", nil
	case int:
		return strconv.Itoa(v), nil
	case uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
		return fmt.Sprintf("%d", v), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'g', -1, 64), nil
	case float64:
		return strconv.FormatFloat(v, 'g', -1, 64), nil
	case fmt.Stringer:
		return v.String(), nil
	case []byte:
		return string(v), nil
	default:
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Pointer || rv.Kind() == reflect.Interface {
			rv = rv.Elem()
			if !rv.IsValid() {
				return "", nil
			}
			return AsString(rv.Interface())
		}
	}
	return "", fmt.Errorf("failed to cast %T to string", value)
}

func asString(value reflect.Value) (reflect.Value, error) {
	r, err := AsString(value.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(r), nil
}
