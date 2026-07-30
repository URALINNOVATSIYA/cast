package cast

import (
	"testing"
)

func TestAsInt8(t *testing.T) {
	str := "42"
	var nilPtr *int8
	tests := []castTest[int8]{
		{nil, 0, ""},
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
		{nilPtr, 0, ""},
		{customAny(nil), 0, ""},
		{customInt(42), 42, ""},
		{customUint(42), 42, ""},
		{customInt8(42), 42, ""},
		{customUint8(42), 42, ""},
		{customInt16(42), 42, ""},
		{customUint16(42), 42, ""},
		{customInt32(42), 42, ""},
		{customUint32(42), 42, ""},
		{customInt64(42), 42, ""},
		{customUint64(42), 42, ""},
		{customFloat32(42.0), 42, ""},
		{customFloat64(42.0), 42, ""},
		{customBool(true), 1, ""},
		{customBool(false), 0, ""},
		{customString(str), 42, ""},
		{"invalid", 0, "failed to cast string to int8"},
		{[]int{1, 2, 3}, 0, "failed to cast []int to int8"},
	}
	runCastTests(t, "AsInt8", AsInt8, tests)
}
