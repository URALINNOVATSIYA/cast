package cast

import (
	"fmt"
	"reflect"
)

func Converter[V any]() (func(any) (V, error), error) {
	var v V
	switch any(v).(type) {
	case int:
		return any(AsInt).(func(any) (V, error)), nil
	case uint:
		return any(AsUint).(func(any) (V, error)), nil
	case int8:
		return any(AsInt8).(func(any) (V, error)), nil
	case uint8:
		return any(AsUint8).(func(any) (V, error)), nil
	case int16:
		return any(AsInt16).(func(any) (V, error)), nil
	case uint16:
		return any(AsUint16).(func(any) (V, error)), nil
	case int32:
		return any(AsInt32).(func(any) (V, error)), nil
	case uint32:
		return any(AsUint32).(func(any) (V, error)), nil
	case int64:
		return any(AsInt64).(func(any) (V, error)), nil
	case uint64:
		return any(AsUint64).(func(any) (V, error)), nil
	case float32:
		return any(AsFloat32).(func(any) (V, error)), nil
	case float64:
		return any(AsFloat64).(func(any) (V, error)), nil
	case string:
		return any(AsString).(func(any) (V, error)), nil
	case bool:
		return any(AsBool).(func(any) (V, error)), nil
	default:
		t := reflect.TypeOf(v)
		switch t.Kind() {
		case reflect.Slice:
			return any(asTypedSlice[V]).(func(any) (V, error)), nil
		case reflect.Map:
			return any(asTypedMap[V]).(func(any) (V, error)), nil
		case reflect.Struct:
			return any(AsStruct[V]).(func(any) (V, error)), nil
		case reflect.Pointer:
			return any(asTypedPointer[V]).(func(any) (V, error)), nil
		case reflect.Interface:
			return any(AsInterface[V]).(func(any) (V, error)), nil
		}
	}
	return nil, fmt.Errorf("unsupported casting to type %T", v)
}

func converter(t reflect.Type) (func(reflect.Value) (reflect.Value, error), error) {
	switch t.Kind() {
	case reflect.Int:
		return asInt, nil
	case reflect.Uint:
		return asUint, nil
	case reflect.Int8:
		return asInt8, nil
	case reflect.Uint8:
		return asUint8, nil
	case reflect.Int16:
		return asInt16, nil
	case reflect.Uint16:
		return asUint16, nil
	case reflect.Int32:
		return asInt32, nil
	case reflect.Uint32:
		return asUint32, nil
	case reflect.Int64:
		return asInt64, nil
	case reflect.Uint64:
		return asUint64, nil
	case reflect.Float32:
		return asFloat32, nil
	case reflect.Float64:
		return asFloat64, nil
	case reflect.String:
		return asString, nil
	case reflect.Bool:
		return asBool, nil
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
