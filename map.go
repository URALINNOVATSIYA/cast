package cast

import (
	"fmt"
	"reflect"
)

func ToMap[K comparable, V any](value any) map[K]V {
	r, err := AsMap[K, V](value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsMap[K comparable, V any](value any) (map[K]V, error) {
	if value == nil {
		return nil, nil
	}
	if v, ok := value.(map[K]V); ok {
		return v, nil
	}

	value = elemOf(reflect.ValueOf(value)).Interface()
	if reflect.TypeOf(value).Kind() != reflect.Map {
		return nil, fmt.Errorf("failed to cast %T to %T", value, map[K]V(nil))
	}

	convertKey, err := Converter[K]()
	if err != nil {
		return nil, err
	}
	convertValue, err := Converter[V]()
	if err != nil {
		return nil, err
	}

	result := make(map[K]V)
	mapValue := reflect.ValueOf(value)
	for _, key := range mapValue.MapKeys() {
		k, err := convertKey(key.Interface())
		if err != nil {
			return nil, err
		}
		v, err := convertValue(mapValue.MapIndex(key).Interface())
		if err != nil {
			return nil, err
		}
		result[k] = v
	}

	return result, nil
}

func asMap(mapType reflect.Type) func(reflect.Value) (reflect.Value, error) {
	return func(value reflect.Value) (reflect.Value, error) {
		value = elemOf(value)
		if !value.IsValid() {
			return reflect.New(mapType).Elem(), nil
		}
		valueType := value.Type()
		if valueType == mapType {
			return value, nil
		}
		if valueType.Kind() != reflect.Map {
			return reflect.Value{}, fmt.Errorf("failed to cast %s to %s", valueType, mapType)
		}
		if value.IsNil() {
			return reflect.New(mapType).Elem(), nil
		}

		mapKeyType := mapType.Key()
		convertKey, err := converter(mapKeyType)
		if err != nil {
			return reflect.Value{}, err
		}
		mapValueType := mapType.Elem()
		convertValue, err := converter(mapValueType)
		if err != nil {
			return reflect.Value{}, err
		}

		result := reflect.MakeMap(mapType)
		for _, key := range value.MapKeys() {
			k, err := convertKey(key)
			if err != nil {
				return reflect.Value{}, err
			}
			v, err := convertValue(value.MapIndex(key))
			if err != nil {
				return reflect.Value{}, err
			}
			result.SetMapIndex(k, v)
		}

		return result, nil
	}
}

func asTypedMap[V any](value any) (V, error) {
	var zero V
	r, err := asMap(reflect.TypeOf(zero))(reflect.ValueOf(value))
	if err != nil {
		return zero, err
	}
	return r.Interface().(V), nil
}
