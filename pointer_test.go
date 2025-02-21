package cast

import (
	"testing"
)

func TestAsPointer(t *testing.T) {
	t.Run("with scalar types", func(t *testing.T) {
		v0 := int64(42)
		v1 := uint8(42)
		v2 := []int{1, 2, 3}
		tests := []castTest[*int64]{
			{nil, nil, ""},
			{int64(42), &v0, ""},
			{int(42), &v0, ""},
			{&v0, &v0, ""},
			{&v1, &v0, ""},
			{&v2, nil, "failed to cast []int to int64"},
			{v2, nil, "failed to cast []int to int64"},
		}
		runCastTests(t, "AsPointer[int64]", AsPointer[int64], tests)
	})

	t.Run("with complex types", func(t *testing.T) {
		v0 := 1
		v1 := 2
		v2 := 3
		v3 := []*int{&v0, &v1, &v2}
		v4 := []int8{1, 2, 3}
		v5 := &v2
		v6 := []any{&v0, v1, &v5}
		v7 := [][]int{{1, 2}, {3}}
		tests := []castTest[*[]*int]{
			{nil, nil, ""},
			{&v3, &v3, ""},
			{v3, &v3, ""},
			{v4, &v3, ""},
			{&v4, &v3, ""},
			{v6, &v3, ""},
			{&v6, &v3, ""},
			{v2, nil, "failed to cast int to []*int"},
			{v7, nil, "failed to cast []int to int"},
		}
		runCastTests(t, "AsPointer[[]*int]", AsPointer[[]*int], tests)
	})
}
