package cast

import (
	"reflect"
	"testing"
	"time"
)

type castTest[V any] struct {
	input    any
	expected V
	err      string
}

func runCastTests[V any](t *testing.T, castName string, convert func(any) (V, error), tests []castTest[V]) {
	for _, test := range tests {
		actual, err := convert(test.input)
		if err == nil {
			if test.err != "" {
				t.Errorf("%s(%#v) must return error %q, got none", castName, test.input, test.err)
			}
			if !equal(actual, test.expected) {
				t.Errorf("%s(%#v) must return %#v, got %#v", castName, test.input, test.expected, actual)
			}
		} else if test.err == "" {
			t.Errorf("%s(%#v) must return no error, got %q", castName, test.input, err)
		} else if test.err != err.Error() {
			t.Errorf("%s(%#v) must return error %q, got %q", castName, test.input, test.err, err)
		}
	}
}

func equal[T any](v1, v2 T) bool {
	switch v := any(v1).(type) {
	case time.Time:
		return v.Equal(any(v2).(time.Time))
	default:
		return reflect.DeepEqual(v1, v2)
	}
}
