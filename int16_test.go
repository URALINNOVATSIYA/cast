package cast

import (
	"testing"
)

func TestAsInt16(t *testing.T) {
	str := "42"
	tests := []castTest[int16]{
		{int(42), 42, ""},
		{uint(42), 42, ""},
		{int8(42), 42, ""},
		{uint8(42), 42, ""},
		{int16(42), 42, ""},
		{uint16(42), 42, ""},
		{int32(42), 42, ""},
		{uint32(42), 42, ""},
		{int64(42), 42, ""},
		{uint64(42), 42, ""},
		{float32(42.0), 42, ""},
		{float64(42.0), 42, ""},
		{true, 1, ""},
		{false, 0, ""},
		{str, 42, ""},
		{&str, 42, ""},
		{nil, 0, ""},
		{"invalid", 0, "strconv.ParseInt: parsing \"invalid\": invalid syntax"},
		{[]int{1, 2, 3}, 0, "failed to cast []int to int16"},
	}
	runCastTests(t, "AsInt16", AsInt16, tests)
}
