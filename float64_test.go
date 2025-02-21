package cast

import (
	"testing"
)

func TestAsFloat64(t *testing.T) {
	str := "42"
	tests := []castTest[float64]{
		{nil, 0, ""},
		{int(42), 42.0, ""},
		{uint(42), 42.0, ""},
		{int8(42), 42.0, ""},
		{uint8(42), 42.0, ""},
		{int16(42), 42.0, ""},
		{uint16(42), 42.0, ""},
		{int32(42), 42.0, ""},
		{uint32(42), 42.0, ""},
		{int64(42), 42.0, ""},
		{uint64(42), 42.0, ""},
		{float32(42), 42.0, ""},
		{float64(42), 42.0, ""},
		{true, 1.0, ""},
		{false, 0.0, ""},
		{str, 42.0, ""},
		{&str, 42.0, ""},
		{"invalid", 0, "strconv.ParseFloat: parsing \"invalid\": invalid syntax"},
		{[]int{1, 2, 3}, 0, "failed to cast []int to float64"},
	}
	runCastTests(t, "AsFloat64", AsFloat64, tests)
}
