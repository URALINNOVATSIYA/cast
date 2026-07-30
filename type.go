package cast

import (
	"fmt"
	"reflect"
)

func ToType[V any](value any) V {
	r, err := AsType[V](value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsType[V any](value any) (V, error) {
	var zero V
	if value == nil {
		return zero, nil
	}
	fn, err := Converter[V]()
	if err != nil {
		return zero, err
	}
	return fn(value)
}

func asType(
	t reflect.Type,
	convert func(reflect.Value) (reflect.Value, error),
) func(reflect.Value) (reflect.Value, error) {
	return func(v reflect.Value) (reflect.Value, error) {
		v, err := convert(v)
		if err != nil {
			return reflect.Value{}, err
		}
		return v.Convert(t), nil
	}
}

func asTypedType[V any](
	t reflect.Type,
	convert func(reflect.Value) (reflect.Value, error),
) func(any) (V, error) {
	return func(value any) (V, error) {
		v, err := convert(reflect.ValueOf(value))
		if err != nil {
			var zero V
			return zero, fmt.Errorf("failed to cast %T to %s", value, t)
		}
		return v.Convert(t).Interface().(V), nil
	}
}
