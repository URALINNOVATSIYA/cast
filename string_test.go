package cast

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestAsString(t *testing.T) {
	v0 := 42
	var nilPtr *string
	uid := uuid.New()
	stringer := &strings.Builder{}
	stringer.WriteString("test")
	tests := []castTest[string]{
		{nil, "", ""},
		{"hello", "hello", ""},
		{true, "true", ""},
		{false, "false", ""},
		{v0, "42", ""},
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
		{customAny(nil), "", ""},
		{customString("hello"), "hello", ""},
		{customBool(true), "true", ""},
		{customBool(false), "false", ""},
		{customInt(v0), "42", ""},
		{customUint(v0), "42", ""},
		{customInt8(42), "42", ""},
		{customUint8(42), "42", ""},
		{customInt16(42), "42", ""},
		{customUint16(42), "42", ""},
		{customInt32(42), "42", ""},
		{customUint32(42), "42", ""},
		{customInt64(42), "42", ""},
		{customUint64(42), "42", ""},
		{customFloat32(42), "42", ""},
		{customFloat32(42.5), "42.5", ""},
		{customFloat64(42), "42", ""},
		{customFloat64(42.5), "42.5", ""},
		{&v0, "42", ""},
		{nilPtr, "", ""},
		{uid, uid.String(), ""},
		{time.Date(2020, 11, 5, 3, 45, 13, 123456, time.UTC), "2020-11-05 03:45:13.000123456 +0000 UTC", ""},
		{customTime(time.Date(2020, 11, 5, 3, 45, 13, 123456, time.UTC)), "2020-11-05 03:45:13.000123456 +0000 UTC", ""},
		{stringer, "test", ""},
		{[]byte{69, 70, 71}, "EFG", ""},
		{customSlice([]byte{50, 51, 52}), "234", ""},
		{[4]byte{49, 50, 51, 52}, "1234", ""},
		{[]int{1, 2, 3}, "", "failed to cast []int to string"},
	}
	runCastTests(t, "AsString", AsString, tests)
}
