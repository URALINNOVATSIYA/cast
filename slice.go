package cast

import (
	"fmt"
	"reflect"
)

func ToSlice[V any](value any) []V {
	r, err := AsSlice[V](value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsSlice[V any](value any) ([]V, error) {
	if value == nil {
		return nil, nil
	}
	if v, ok := value.([]V); ok {
		return v, nil
	}

	value = elemOf(reflect.ValueOf(value)).Interface()
	if reflect.TypeOf(value).Kind() != reflect.Slice {
		return nil, fmt.Errorf("failed to cast %T to %T", value, []V(nil))
	}

	convert, err := Converter[V]()
	if err != nil {
		return nil, err
	}

	slice := reflect.ValueOf(value)
	size := slice.Len()
	result := make([]V, size)
	for i := 0; i < size; i++ {
		item := slice.Index(i).Interface()
		v, err := convert(item)
		if err != nil {
			return nil, err
		}
		result[i] = v
	}

	return result, nil
}

func asSlice(sliceType reflect.Type) func(reflect.Value) (reflect.Value, error) {
	return func(value reflect.Value) (reflect.Value, error) {
		value = elemOf(value)
		if !value.IsValid() {
			return reflect.New(sliceType).Elem(), nil
		}
		valueType := value.Type()
		if valueType == sliceType {
			return value, nil
		}
		if valueType.Kind() != reflect.Slice {
			return reflect.Value{}, fmt.Errorf("failed to cast %s to %s", valueType, sliceType)
		}
		if value.IsNil() {
			return reflect.New(sliceType).Elem(), nil
		}

		sliceElementType := sliceType.Elem()
		convert, err := converter(sliceElementType)
		if err != nil {
			return reflect.Value{}, err
		}

		size := value.Len()
		result := reflect.MakeSlice(sliceType, size, value.Cap())
		for i := 0; i < size; i++ {
			v, err := convert(value.Index(i))
			if err != nil {
				return reflect.Value{}, err
			}
			result.Index(i).Set(v)
		}

		return result, nil
	}
}

func asTypedSlice[V any](value any) (V, error) {
	var zero V
	r, err := asSlice(reflect.TypeOf(zero))(reflect.ValueOf(value))
	if err != nil {
		return zero, err
	}
	return r.Interface().(V), nil
}
