package cast

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
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
		switch rv.Kind() {
		case reflect.Bool, reflect.String,
			reflect.Int, reflect.Uint,
			reflect.Int8, reflect.Uint8,
			reflect.Int16, reflect.Uint16,
			reflect.Int32, reflect.Uint32,
			reflect.Int64, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			return AsString(rv.Convert(kind2types[rv.Kind()]).Interface())
		case reflect.Slice:
			if rv.Type().Elem().Kind() == reflect.Uint8 {
				return AsString(rv.Convert(typeByteSlice).Interface())
			}
		case reflect.Array:
			if rv.Type().Elem().Kind() == reflect.Uint8 {
				if rv.CanAddr() {
					return AsString(rv.Slice(0, rv.Len()).Interface())
				}
				p := reflect.New(rv.Type())
				p.Elem().Set(rv)
				return AsString(p.Elem().Slice(0, rv.Len()).Interface())
			}
		case reflect.Pointer, reflect.Interface:
			rv = rv.Elem()
			if !rv.IsValid() {
				return "", nil
			}
			return AsString(rv.Interface())
		case reflect.Struct:
			if rv.CanConvert(typeTime) {
				return rv.Convert(typeTime).Interface().(time.Time).String(), nil
			}
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
