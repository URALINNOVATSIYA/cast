package cast

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

func ToUuid(value any) uuid.UUID {
	r, err := AsUuid(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsUuid(value any) (uuid.UUID, error) {
	switch v := value.(type) {
	case nil:
		return uuid.Nil, nil
	case uuid.UUID:
		return v, nil
	case string:
		return uuid.Parse(v)
	case []byte:
		if len(v) == 16 {
			return uuid.FromBytes(v)
		}
		return uuid.ParseBytes(v)
	default:
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.String:
			return AsUuid(rv.Convert(kind2types[rv.Kind()]).Interface())
		case reflect.Slice:
			if rv.Type().Elem().Kind() == reflect.Uint8 {
				return AsUuid(rv.Convert(typeByteSlice).Interface())
			}
		case reflect.Array:
			if rv.Type().Elem().Kind() == reflect.Uint8 && rv.Len() == 16 {
				return AsUuid(rv.Convert(typeUUID).Interface())
			}
		case reflect.Pointer, reflect.Interface:
			rv = rv.Elem()
			if !rv.IsValid() {
				return uuid.Nil, nil
			}
			return AsUuid(rv.Interface())
		}
	}
	return uuid.Nil, fmt.Errorf("failed to cast %T to UUID", value)
}
