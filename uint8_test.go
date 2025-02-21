package cast

import (
	"testing"
)

func TestAsUint8(t *testing.T) {
	str := "42"
	tests := []castTest[uint8]{
		{int(42), 42, ""},
		{uint(42), 42, ""},
		{int8(42), 42, ""},
		{uint8(42), 42, ""},
		{int16(42), 42, ""},
		{uint8(42), 42, ""},
		{int32(42), 42, ""},
		{uint8(42), 42, ""},
		{int64(42), 42, ""},
		{uint64(42), 42, ""},
		{float32(42.0), 42, ""},
		{float64(42.0), 42, ""},
		{true, 1, ""},
		{false, 0, ""},
		{str, 42, ""},
		{&str, 42, ""},
		{nil, 0, ""},
		{"invalid", 0, "strconv.ParseUint: parsing \"invalid\": invalid syntax"},
		{[]int{1, 2, 3}, 0, "failed to cast []int to uint8"},
	}
	runCastTests(t, "AsUint8", AsUint8, tests)
}
