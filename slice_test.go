package cast

import (
	"testing"
)

func TestAsSlice(t *testing.T) {
	t.Run("with scalar element type", func(t *testing.T) {
		v0 := []string{"1", "2", "3"}
		expected := []int{1, 2, 3}
		tests := []castTest[[]int]{
			{nil, nil, ""},
			{[]int{}, []int{}, ""},
			{[]string{}, []int{}, ""},
			{[]int{1, 2, 3}, expected, ""},
			{[]uint{1, 2, 3}, expected, ""},
			{[]int8{1, 2, 3}, expected, ""},
			{[]uint8{1, 2, 3}, expected, ""},
			{[]int16{1, 2, 3}, expected, ""},
			{[]uint16{1, 2, 3}, expected, ""},
			{[]int32{1, 2, 3}, expected, ""},
			{[]uint32{1, 2, 3}, expected, ""},
			{[]int64{1, 2, 3}, expected, ""},
			{[]uint64{1, 2, 3}, expected, ""},
			{[]float32{1, 2, 3}, expected, ""},
			{[]float64{1, 2, 3}, expected, ""},
			{[]bool{true, true, false}, []int{1, 1, 0}, ""},
			{v0, expected, ""},
			{&v0, expected, ""},
			{"invalid", nil, "failed to cast string to []int"},
		}
		runCastTests(t, "AsSlice[int]", AsSlice[int], tests)
	})

	t.Run("with complex element type", func(t *testing.T) {
		v0 := [][]string{{"1", "2"}, {"3"}, nil}
		expected := [][]string{{"1", "2"}, {"3"}, nil}
		floatExpected := [][]string{{"1", "2"}, {"3"}, nil}
		tests := []castTest[[][]string]{
			{nil, nil, ""},
			{[][]int{}, [][]string{}, ""},
			{[][]bool{}, [][]string{}, ""},
			{[][]int{{1, 2}, {3}, nil}, expected, ""},
			{[][]uint{{1, 2}, {3}, nil}, expected, ""},
			{[][]int8{{1, 2}, {3}, nil}, expected, ""},
			{[][]uint8{{1, 2}, {3}, nil}, expected, ""},
			{[][]int16{{1, 2}, {3}, nil}, expected, ""},
			{[][]uint16{{1, 2}, {3}, nil}, expected, ""},
			{[][]int32{{1, 2}, {3}, nil}, expected, ""},
			{[][]uint32{{1, 2}, {3}, nil}, expected, ""},
			{[][]int64{{1, 2}, {3}, nil}, expected, ""},
			{[][]uint64{{1, 2}, {3}, nil}, expected, ""},
			{[][]float32{{1, 2}, {3}, nil}, floatExpected, ""},
			{[][]float64{{1, 2}, {3}, nil}, floatExpected, ""},
			{[][]bool{{true, true}, {false}, nil}, [][]string{{"true", "true"}, {"false"}, nil}, ""},
			{v0, expected, ""},
			{&v0, expected, ""},
			{[]string{"invalid"}, nil, "failed to cast string to []string"},
			{123, nil, "failed to cast int to [][]string"},
		}
		runCastTests(t, "AsSlice[[]string]", AsSlice[[]string], tests)
	})
}
