package cast

import (
	"testing"
)

func TestAsBool(t *testing.T) {
	str := "true"
	tests := []castTest[bool]{
		{nil, false, ""},
		{true, true, ""},
		{false, false, ""},
		{uint(1), true, ""},
		{uint(0), false, ""},
		{int(1), true, ""},
		{int(0), false, ""},
		{int8(1), true, ""},
		{int8(0), false, ""},
		{uint8(1), true, ""},
		{uint8(0), false, ""},
		{int16(1), true, ""},
		{int16(0), false, ""},
		{uint16(1), true, ""},
		{uint16(0), false, ""},
		{int32(1), true, ""},
		{int32(0), false, ""},
		{uint32(1), true, ""},
		{uint32(0), false, ""},
		{int64(1), true, ""},
		{int64(0), false, ""},
		{uint64(1), true, ""},
		{uint64(0), false, ""},
		{float32(0), false, ""},
		{float32(1), true, ""},
		{float64(0), false, ""},
		{float64(1), true, ""},
		{"true", true, ""},
		{"false", false, ""},
		{&str, true, ""},
		{"foo", false, "failed to cast \"foo\" to bool"},
		{struct{}{}, false, "failed to cast struct {} to bool"},
	}
	runCastTests(t, "AsBool", AsBool, tests)
}
