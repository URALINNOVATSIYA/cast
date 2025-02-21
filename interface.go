package cast

import (
	"fmt"
	"reflect"
)

func ToInterface[V any](value any) V {
	r, err := AsInterface[V](value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsInterface[V any](value any) (V, error) {
	var result V
	if value == nil {
		return result, nil
	}
	if v, ok := value.(V); ok {
		return v, nil
	}
	convert := asInterface(reflect.TypeOf(&result).Elem())
	r, err := convert(reflect.ValueOf(value))
	if err != nil {
		return result, err
	}
	return r.Interface().(V), nil
}

func asInterface(interfaceType reflect.Type) func(reflect.Value) (reflect.Value, error) {
	return func(value reflect.Value) (reflect.Value, error) {
		valueType := value.Type()
		if valueType == interfaceType {
			return value, nil
		}
		if !value.CanConvert(interfaceType) {
			return reflect.Value{}, fmt.Errorf("failed to cast %s to %s", valueType, interfaceType)
		}
		return value.Convert(interfaceType), nil
	}
}
