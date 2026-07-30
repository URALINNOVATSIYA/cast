package cast

import (
	"testing"
)

func TestAsFloat32(t *testing.T) {
	str := "42"
	var nilPtr *float32
	tests := []castTest[float32]{
		{nil, 0, ""},
		{uint(42), 42.0, ""},
		{int(42), 42.0, ""},
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
		{nilPtr, 0, ""},
		{customAny(nil), 0, ""},
		{customUint(42), 42.0, ""},
		{customInt(42), 42.0, ""},
		{customInt8(42), 42.0, ""},
		{customUint8(42), 42.0, ""},
		{customInt16(42), 42.0, ""},
		{customUint16(42), 42.0, ""},
		{customInt32(42), 42.0, ""},
		{customUint32(42), 42.0, ""},
		{customInt64(42), 42.0, ""},
		{customUint64(42), 42.0, ""},
		{customFloat32(3.45), 3.45, ""},
		{customFloat64(3.45), 3.45, ""},
		{customBool(true), 1.0, ""},
		{customBool(false), 0.0, ""},
		{customString(str), 42.0, ""},
		{"invalid", 0, "strconv.ParseFloat: parsing \"invalid\": invalid syntax"},
		{[]int{1, 2, 3}, 0, "failed to cast []int to float32"},
	}
	runCastTests(t, "AsFloat32", AsFloat32, tests)
}
