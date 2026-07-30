package cast

import (
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
)

func Converter[V any, C func(any) (V, error)]() (C, error) {
	var v V
	switch any(v).(type) {
	case nil:
		return any(AsInterface[V]).(C), nil
	case int:
		return any(AsInt).(C), nil
	case uint:
		return any(AsUint).(C), nil
	case int8:
		return any(AsInt8).(C), nil
	case uint8:
		return any(AsUint8).(C), nil
	case int16:
		return any(AsInt16).(C), nil
	case uint16:
		return any(AsUint16).(C), nil
	case int32:
		return any(AsInt32).(C), nil
	case uint32:
		return any(AsUint32).(C), nil
	case int64:
		return any(AsInt64).(C), nil
	case uint64:
		return any(AsUint64).(C), nil
	case float32:
		return any(AsFloat32).(C), nil
	case float64:
		return any(AsFloat64).(C), nil
	case string:
		return any(AsString).(C), nil
	case bool:
		return any(AsBool).(C), nil
	case uuid.UUID:
		return any(AsUuid).(C), nil
	case time.Time:
		return any(AsTime).(C), nil
	default:
		t := reflect.TypeOf(v)
		switch t.Kind() {
		case reflect.Int:
			return asTypedType[V](t, asInt), nil
		case reflect.Uint:
			return asTypedType[V](t, asUint), nil
		case reflect.Int8:
			return asTypedType[V](t, asInt8), nil
		case reflect.Uint8:
			return asTypedType[V](t, asUint8), nil
		case reflect.Int16:
			return asTypedType[V](t, asInt16), nil
		case reflect.Uint16:
			return asTypedType[V](t, asUint16), nil
		case reflect.Int32:
			return asTypedType[V](t, asInt32), nil
		case reflect.Uint32:
			return asTypedType[V](t, asUint32), nil
		case reflect.Int64:
			return asTypedType[V](t, asInt64), nil
		case reflect.Uint64:
			return asTypedType[V](t, asUint64), nil
		case reflect.Float32:
			return asTypedType[V](t, asFloat32), nil
		case reflect.Float64:
			return asTypedType[V](t, asFloat64), nil
		case reflect.String:
			return asTypedType[V](t, asString), nil
		case reflect.Bool:
			return asTypedType[V](t, asBool), nil
		case reflect.Array:
			return asTypedArray[V], nil
		case reflect.Slice:
			return asTypedSlice[V], nil
		case reflect.Map:
			return asTypedMap[V], nil
		case reflect.Struct:
			return AsStruct[V], nil
		case reflect.Pointer:
			return asTypedPointer[V], nil
		case reflect.Interface:
			return AsInterface[V], nil
		}
	}
	return nil, fmt.Errorf("unsupported casting to type %T", v)
}

func converter(t reflect.Type) (func(reflect.Value) (reflect.Value, error), error) {
	switch t.Kind() {
	case reflect.Int:
		return asType(t, asInt), nil
	case reflect.Uint:
		return asType(t, asUint), nil
	case reflect.Int8:
		return asType(t, asInt8), nil
	case reflect.Uint8:
		return asType(t, asUint8), nil
	case reflect.Int16:
		return asType(t, asInt16), nil
	case reflect.Uint16:
		return asType(t, asUint16), nil
	case reflect.Int32:
		return asType(t, asInt32), nil
	case reflect.Uint32:
		return asType(t, asUint32), nil
	case reflect.Int64:
		return asType(t, asInt64), nil
	case reflect.Uint64:
		return asType(t, asUint64), nil
	case reflect.Float32:
		return asType(t, asFloat32), nil
	case reflect.Float64:
		return asType(t, asFloat64), nil
	case reflect.String:
		return asType(t, asString), nil
	case reflect.Bool:
		return asType(t, asBool), nil
	case reflect.Array:
		return asArray(t), nil
	case reflect.Slice:
		return asSlice(t), nil
	case reflect.Map:
		return asMap(t), nil
	case reflect.Struct:
		return asStruct(t), nil
	case reflect.Pointer:
		return asPointer(t), nil
	case reflect.Interface:
		return asInterface(t), nil
	}
	return nil, fmt.Errorf("unsupported casting to type %s", t)
}
