package cast

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
