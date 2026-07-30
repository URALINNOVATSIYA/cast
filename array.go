package cast

import (
	"fmt"
	"reflect"
)

type Array1[T any] interface {
	~[1]T
}

type Array2[T any] interface {
	~[2]T
}

type Array3[T any] interface {
	~[3]T
}

type Array4[T any] interface {
	~[4]T
}

type Array5[T any] interface {
	~[5]T
}

type Array6[T any] interface {
	~[6]T
}

type Array7[T any] interface {
	~[7]T
}

type Array8[T any] interface {
	~[8]T
}

type Array9[T any] interface {
	~[9]T
}

type Array10[T any] interface {
	~[10]T
}

func ToArray1[T any, A Array1[T]](value any) A {
	return toArray[A](value)
}

func ToArray2[T any, A Array2[T]](value any) A {
	return toArray[A](value)
}

func ToArray3[T any, A Array3[T]](value any) A {
	return toArray[A](value)
}

func ToArray4[T any, A Array4[T]](value any) A {
	return toArray[A](value)
}

func ToArray5[T any, A Array5[T]](value any) A {
	return toArray[A](value)
}

func ToArray6[T any, A Array6[T]](value any) A {
	return toArray[A](value)
}

func ToArray7[T any, A Array7[T]](value any) A {
	return toArray[A](value)
}

func ToArray8[T any, A Array8[T]](value any) A {
	return toArray[A](value)
}

func ToArray9[T any, A Array9[T]](value any) A {
	return toArray[A](value)
}

func ToArray10[T any, A Array10[T]](value any) A {
	return toArray[A](value)
}

func toArray[A any](value any) A {
	r, err := asTypedArray[A](value)
	if err != nil {
		panic(err)
	}
	return r
}

func ToArray[A any](value any) A {
	r, err := AsArray[A](value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsArray1[T any, A Array1[T]](value any) (A, error) {
	return asTypedArray[A](value)
}

func AsArray2[T any, A Array2[T]](value any) (A, error) {
	return asTypedArray[A](value)
}

func AsArray3[T any, A Array3[T]](value any) (A, error) {
	return asTypedArray[A](value)
}

func AsArray4[T any, A Array4[T]](value any) (A, error) {
	return asTypedArray[A](value)
}

func AsArray5[T any, A Array5[T]](value any) (A, error) {
	return asTypedArray[A](value)
}

func AsArray6[T any, A Array6[T]](value any) (A, error) {
	return asTypedArray[A](value)
}

func AsArray7[T any, A Array7[T]](value any) (A, error) {
	return asTypedArray[A](value)
}

func AsArray8[T any, A Array8[T]](value any) (A, error) {
	return asTypedArray[A](value)
}

func AsArray9[T any, A Array9[T]](value any) (A, error) {
	return asTypedArray[A](value)
}

func AsArray10[T any, A Array10[T]](value any) (A, error) {
	return asTypedArray[A](value)
}

func AsArray[A any](value any) (A, error) {
	if reflect.TypeFor[A]().Kind() != reflect.Array {
		var zero A
		return zero, fmt.Errorf("generic type must be an array type, %T given", zero)
	}
	return asTypedArray[A](value)
}

func asArray(arrayType reflect.Type) func(reflect.Value) (reflect.Value, error) {
	return func(value reflect.Value) (reflect.Value, error) {
		value = elemOf(value)
		if !value.IsValid() {
			return reflect.New(arrayType).Elem(), nil
		}
		valueType := value.Type()
		if valueType == arrayType {
			return value, nil
		}
		switch valueType.Kind() {
		case reflect.Slice, reflect.Array, reflect.String:
			arrayElementType := arrayType.Elem()
			convert, err := converter(arrayElementType)
			if err != nil {
				return reflect.Value{}, err
			}
			size := value.Len()
			if size == arrayType.Len() {
				result := reflect.New(reflect.ArrayOf(size, arrayElementType)).Elem()
				for i := range size {
					v, err := convert(value.Index(i))
					if err != nil {
						return reflect.Value{}, err
					}
					result.Index(i).Set(v)
				}
				return result.Convert(arrayType), nil
			}
		}
		return reflect.Value{}, fmt.Errorf("failed to cast %s to %s", valueType, arrayType)
	}
}

func asTypedArray[V any](value any) (V, error) {
	var zero V
	r, err := asArray(reflect.TypeOf(zero))(reflect.ValueOf(value))
	if err != nil {
		return zero, err
	}
	return r.Interface().(V), nil
}
