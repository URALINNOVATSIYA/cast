package cast

import (
	"fmt"
	"reflect"
)

func ToStruct[S any](value any) S {
	r, err := AsStruct[S](value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsStruct[S any](value any) (S, error) {
	var result S
	if value == nil {
		return result, nil
	}
	if v, ok := value.(S); ok {
		return v, nil
	}

	rs := reflect.ValueOf(&result).Elem()
	if rs.Kind() != reflect.Struct {
		return result, fmt.Errorf("expected generic type must be struct, %T got", result)
	}

	rv := elemOf(reflect.ValueOf(value))
	switch rv.Kind() {
	case reflect.Map:
		r, err := mapAsStruct(rv, rs)
		if err != nil {
			return result, err
		}
		return r.Interface().(S), err
	case reflect.Struct:
		r, err := structAsStruct(rv, rs)
		if err != nil {
			return result, err
		}
		return r.Interface().(S), err
	}
	return result, fmt.Errorf("failed to cast %T to %T", value, result)
}

func asStruct(structType reflect.Type) func(reflect.Value) (reflect.Value, error) {
	return func(value reflect.Value) (reflect.Value, error) {
		value = elemOf(value)
		valueType := value.Type()
		if valueType == structType {
			return value, nil
		}
		s := reflect.New(structType).Elem()
		switch value.Kind() {
		case reflect.Map:
			r, err := mapAsStruct(value, s)
			if err != nil {
				return reflect.Value{}, err
			}
			return r, err
		case reflect.Struct:
			r, err := structAsStruct(value, s)
			if err != nil {
				return reflect.Value{}, err
			}
			return r, err
		}
		return reflect.Value{}, fmt.Errorf("failed to cast %s to %s", valueType, structType)
	}
}

func mapAsStruct(m reflect.Value, s reflect.Value) (reflect.Value, error) {
	convertKey, _ := Converter[string]()
	for _, k := range m.MapKeys() {
		key, err := convertKey(k.Interface())
		if err != nil {
			return reflect.Value{}, err
		}
		field := s.FieldByName(key)
		if !field.IsValid() {
			field = s.FieldByName(ucfirst(key))
			if !field.IsValid() {
				field = s.FieldByName(lcfirst(key))
				if !field.IsValid() {
					continue
				}
			}
		}
		if !field.CanInterface() {
			continue
		}
		convertValue, err := converter(field.Type())
		if err != nil {
			return reflect.Value{}, err
		}
		value, err := convertValue(m.MapIndex(k))
		if err != nil {
			return reflect.Value{}, err
		}
		field.Set(value)
	}
	return s, nil
}

func structAsStruct(from reflect.Value, to reflect.Value) (reflect.Value, error) {
	srcStructType := from.Type()
	for i := from.NumField() - 1; i >= 0; i-- {
		srcStructField := srcStructType.Field(i)
		if !srcStructField.IsExported() {
			continue
		}
		fieldName := srcStructField.Name
		field := to.FieldByName(fieldName)
		if !field.IsValid() {
			field = to.FieldByName(lcfirst(fieldName))
			if !field.IsValid() {
				continue
			}
		}
		if !field.CanInterface() {
			continue
		}
		convertValue, err := converter(field.Type())
		if err != nil {
			return reflect.Value{}, err
		}
		value, err := convertValue(from.FieldByName(fieldName))
		if err != nil {
			return reflect.Value{}, err
		}
		field.Set(value)
	}
	return to, nil
}
