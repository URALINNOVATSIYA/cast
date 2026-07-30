package cast

import "testing"

func TestAsArray(t *testing.T) {
	t.Run("with scalar element type", func(t *testing.T) {
		var nilPtr *[]int
		expected := [3]int{1, 2, 3}
		tests := []castTest[[3]int]{
			{nil, [3]int{}, ""},
			{nilPtr, [3]int{}, ""},
			{&expected, expected, ""},
			{[]int{1, 2, 3}, expected, ""},
			{[]uint{1, 2, 3}, expected, ""},
			{[]int8{1, 2, 3}, expected, ""},
			{[3]int{1, 2, 3}, expected, ""},
			{[3]uint{1, 2, 3}, expected, ""},
			{[]bool{true, true, false}, [3]int{1, 1, 0}, ""},
			{"123", [3]int{49, 50, 51}, ""},
			{"invalid", [3]int{}, "failed to cast string to [3]int"},
			{123, [3]int{}, "failed to cast int to [3]int"},
		}
		runCastTests(t, "AsArray[[3]int]", AsArray[[3]int], tests)
	})

	t.Run("with complex element type", func(t *testing.T) {
		expected := [2][]string{{"1", "2"}, {"3"}}
		tests := []castTest[[2][]string]{
			{nil, [2][]string{}, ""},
			{[][]string{{"1", "2"}, {"3"}}, expected, ""},
			{[2][]string{{"1", "2"}, {"3"}}, expected, ""},
			{customArray([2][]string{{"1", "2"}, {"3"}}), expected, ""},
		}
		runCastTests(t, "AsArray[[2][]string]", AsArray[[2][]string], tests)
	})

	t.Run("with custom array target type", func(t *testing.T) {
		expected := [2][]string{{"1", "2"}, {"3"}}
		tests := []castTest[customArray]{
			{nil, [2][]string{}, ""},
			{[][]string{{"1", "2"}, {"3"}}, expected, ""},
			{[2][]string{{"1", "2"}, {"3"}}, expected, ""},
			{customArray([2][]string{{"1", "2"}, {"3"}}), expected, ""},
		}
		runCastTests(t, "AsArray[customArray]", AsArray[customArray], tests)
	})

	t.Run("with non-array target type", func(t *testing.T) {
		tests := []castTest[int]{
			{nil, 0, "generic type must be an array type, int given"},
			{"invalid", 0, "generic type must be an array type, int given"},
		}
		runCastTests(t, "AsArray[int]", AsArray[int], tests)
	})
}
