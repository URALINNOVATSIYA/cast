package cast

import (
	"testing"
)

func TestAsString(t *testing.T) {
	v0 := 42
	tests := []castTest[string]{
		{nil, "", ""},
		{"hello", "hello", ""},
		{true, "true", ""},
		{false, "false", ""},
		{v0, "42", ""},
		{&v0, "42", ""},
		{int8(42), "42", ""},
		{uint8(42), "42", ""},
		{int16(42), "42", ""},
		{uint16(42), "42", ""},
		{int32(42), "42", ""},
		{uint32(42), "42", ""},
		{int64(42), "42", ""},
		{uint64(42), "42", ""},
		{float32(42), "42", ""},
		{float32(42.5), "42.5", ""},
		{float64(42), "42", ""},
		{float64(42.5), "42.5", ""},
		{[]int{1, 2, 3}, "", "failed to cast []int to string"},
	}
	runCastTests(t, "AsString", AsString, tests)
}
