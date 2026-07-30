package cast

import (
	"fmt"
	"reflect"
	"time"
)

var TimeLayouts = []string{
	time.Kitchen,
	time.DateOnly,
	time.TimeOnly,
	time.DateTime,
	time.RFC3339,
	"2006-01-02 15:04:05.999999-07",
	"2006-01-02 15:04:05.999999-07:00",
	"2006-01-02T15:04:05.999999-07:00",
	"2006-01-02T15:04:05.999999Z",
	"2006-01-02 15:04:05.999",
	"2006-01-02 15:04:05.999-07",
	"2006-01-02 15:04:05.999-07:00",
	"2006-01-02T15:04:05.999-07:00",
	"2006-01-02T15:04:05.999Z",
	time.RFC3339Nano,
	time.UnixDate,
	time.RubyDate,
	time.ANSIC,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	"2006-01-02 15:04:05.999999999 -0700 MST",
}

func ToTime(value any) time.Time {
	r, err := AsTime(value)
	if err != nil {
		panic(err)
	}
	return r
}

func AsTime(value any) (time.Time, error) {
	var t time.Time
	if t, ok := value.(time.Time); ok {
		return t, nil
	}
	switch v := value.(type) {
	case time.Time:
		return v, nil
	default:
		rv := reflect.ValueOf(value)
		if rv.Kind() == reflect.Struct && rv.CanConvert(typeTime) {
			return rv.Convert(typeTime).Interface().(time.Time), nil
		}
	}
	v, err := AsString(value)
	if err != nil {
		return time.Time{}, err
	}
	for _, layout := range TimeLayouts {
		if t, err = time.Parse(layout, v); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("failed to parse %q to time", v)
}
