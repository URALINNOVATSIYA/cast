package cast

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

type customAny any
type customBool bool
type customInt int
type customUint uint
type customInt8 int8
type customUint8 uint8
type customInt16 int16
type customUint16 uint16
type customInt32 int32
type customUint32 uint32
type customInt64 int64
type customUint64 uint64
type customFloat32 float32
type customFloat64 float64
type customString string
type customSlice []byte
type customArray [2][]string
type customMap map[string]any
type customUuid uuid.UUID
type customTime time.Time
type customPointer *customUuid

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

func TestAsType(t *testing.T) {
	t.Run("customBool", func(t *testing.T) {
		str := "true"
		var nilPtr *bool
		tests := []castTest[customBool]{
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
			{nilPtr, false, ""},
			{customAny(nil), false, ""},
			{customBool(true), true, ""},
			{customBool(false), false, ""},
			{customUint(1), true, ""},
			{customUint(0), false, ""},
			{customInt(1), true, ""},
			{customInt(0), false, ""},
			{customInt8(1), true, ""},
			{customInt8(0), false, ""},
			{customUint8(1), true, ""},
			{customUint8(0), false, ""},
			{customInt16(1), true, ""},
			{customInt16(0), false, ""},
			{customUint16(1), true, ""},
			{customUint16(0), false, ""},
			{customInt32(1), true, ""},
			{customInt32(0), false, ""},
			{customUint32(1), true, ""},
			{customUint32(0), false, ""},
			{customInt64(1), true, ""},
			{customInt64(0), false, ""},
			{customUint64(1), true, ""},
			{customUint64(0), false, ""},
			{customFloat32(0), false, ""},
			{customFloat32(1), true, ""},
			{customFloat64(0), false, ""},
			{customFloat64(1), true, ""},
			{customString("true"), true, ""},
			{customString("false"), false, ""},
			{"foo", false, "failed to cast string to cast.customBool"},
			{struct{}{}, false, "failed to cast struct {} to cast.customBool"},
		}
		runCastTests(t, "AsType[customBool]", AsType, tests)
	})

	t.Run("customInt", func(t *testing.T) {
		str := "42"
		var nilPtr *int
		tests := []castTest[customInt]{
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
			{"invalid", 0, "failed to cast string to cast.customInt"},
			{[]int{1, 2, 3}, 0, "failed to cast []int to cast.customInt"},
		}
		runCastTests(t, "AsType[customInt]", AsType, tests)
	})

	t.Run("customString", func(t *testing.T) {
		v0 := 42
		var nilPtr *string
		uid := uuid.New()
		stringer := &strings.Builder{}
		stringer.WriteString("test")
		tests := []castTest[customString]{
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
			{uid, customString(uid.String()), ""},
			{time.Date(2020, 11, 5, 3, 45, 13, 123456, time.UTC), "2020-11-05 03:45:13.000123456 +0000 UTC", ""},
			{customTime(time.Date(2020, 11, 5, 3, 45, 13, 123456, time.UTC)), "2020-11-05 03:45:13.000123456 +0000 UTC", ""},
			{stringer, "test", ""},
			{[]byte{69, 70, 71}, "EFG", ""},
			{customSlice([]byte{50, 51, 52}), "234", ""},
			{[4]byte{49, 50, 51, 52}, "1234", ""},
			{[]int{1, 2, 3}, "", "failed to cast []int to cast.customString"},
		}
		runCastTests(t, "AsType[customString]", AsType, tests)
	})

	t.Run("customTime", func(t *testing.T) {
		tests := []castTest[customTime]{
			{time.Date(2020, 5, 20, 13, 15, 47, 0, time.UTC), customTime(time.Date(2020, 5, 20, 13, 15, 47, 0, time.UTC)), ""},
			{customString("2025-05-04"), customTime(time.Date(2025, 5, 4, 0, 0, 0, 0, time.UTC)), ""},
			{customTime(time.Date(2020, 5, 20, 13, 15, 47, 0, time.UTC)), customTime(time.Date(2020, 5, 20, 13, 15, 47, 0, time.UTC)), ""},
			{"5:4", customTime(time.Time{}), "failed to parse \"5:4\" to time"},
		}
		runCastTests(t, "AsType[customTime]", AsType, tests)
	})

	t.Run("customPointer", func(t *testing.T) {
		uid := uuid.New()
		cuid := customUuid(uid)
		tests := []castTest[customPointer]{
			{nil, nil, ""},
			{&uid, customPointer(&cuid), ""},
		}
		runCastTests(t, "AsType[customPointer]", AsType, tests)
	})
}
