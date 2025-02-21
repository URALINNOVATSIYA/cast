package cast

import (
	"reflect"
	"unsafe"
)

func ToPointer[V any](value any) *V {
	r, err := AsPointer[V](value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsPointer[V any](value any) (*V, error) {
	if value == nil {
		return nil, nil
	}
	if v, ok := value.(*V); ok {
		return v, nil
	}
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Pointer {
		value = elemOf(rv).Interface()
	}
	fn, err := Converter[V]()
	if err != nil {
		return nil, err
	}
	var v V
	v, err = fn(value)
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func asPointer(pointerType reflect.Type) func(reflect.Value) (reflect.Value, error) {
	return func(value reflect.Value) (reflect.Value, error) {
		valueType := value.Type()
		if valueType == pointerType {
			return value, nil
		}
		switch valueType.Kind() {
		case reflect.Interface, reflect.Pointer, reflect.Map, reflect.Slice:
			if value.IsNil() {
				return reflect.New(pointerType).Elem(), nil
			}
			value = elemOf(value)
		}
		pointerElemType := pointerType.Elem()
		convert, err := converter(pointerElemType)
		if err != nil {
			return reflect.Value{}, err
		}
		v, err := convert(value)
		if err != nil {
			return reflect.Value{}, err
		}
		if v.CanAddr() {
			return reflect.NewAt(pointerType, unsafe.Pointer(v.UnsafeAddr())), nil
		}
		p := reflect.New(v.Type())
		p.Elem().Set(v)
		return p, nil
	}
}

func elemOf(value reflect.Value) reflect.Value {
	for value.Kind() == reflect.Pointer || value.Kind() == reflect.Interface {
		value = value.Elem()
	}
	return value
}

func asTypedPointer[V any](value any) (V, error) {
	var zero V
	r, err := asPointer(reflect.TypeOf(zero))(reflect.ValueOf(value))
	if err != nil {
		return zero, err
	}
	return r.Interface().(V), nil
}
