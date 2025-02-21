package cast

import (
	"testing"
)

func TestAsMap(t *testing.T) {
	t.Run("with scalar types", func(t *testing.T) {
		v0 := map[string]int{"1": 1, "2": 0, "3": 3}
		expected := map[int]bool{1: true, 2: false, 3: true}
		tests := []castTest[map[int]bool]{
			{nil, nil, ""},
			{map[string]int{}, map[int]bool{}, ""},
			{map[int]bool{}, map[int]bool{}, ""},
			{map[int]string{1: "true", 2: "", 3: "true"}, expected, ""},
			{map[uint]bool{1: true, 2: false, 3: true}, expected, ""},
			{map[int8]float64{1: 1.5, 2: 0, 3: -2.3}, expected, ""},
			{map[uint8]float32{1: 1, 2: 0, 3: 3}, expected, ""},
			{map[int16]int8{1: 1, 2: 0, 3: 3}, expected, ""},
			{map[uint16]uint64{1: 1, 2: 0, 3: 3}, expected, ""},
			{map[int32]int64{1: 1, 2: 0, 3: 3}, expected, ""},
			{map[uint32]uint32{1: 1, 2: 0, 3: 3}, expected, ""},
			{map[int64]int32{1: 1, 2: 0, 3: 3}, expected, ""},
			{map[uint64]uint16{1: 1, 2: 0, 3: 3}, expected, ""},
			{map[float32]uint8{1: 1, 2: 0, 3: 3}, expected, ""},
			{map[float64]int8{1: 1, 2: 0, 3: 3}, expected, ""},
			{map[bool]uint{true: 0, false: 1}, map[int]bool{1: false, 0: true}, ""},
			{v0, expected, ""},
			{&v0, expected, ""},
			{"invalid", nil, "failed to cast string to map[int]bool"},
		}
		runCastTests(t, "AsMap[int, bool]", AsMap[int, bool], tests)
	})

	t.Run("with complex types", func(t *testing.T) {
		expected := map[int]map[string]bool{1: {"1": true, "2": false}, 2: {"3": true}, 3: nil}
		floatExpected := map[int]map[string]bool{1: {"1": true, "2": false}, 2: {"3": true}, 3: nil}
		v0 := map[uint64]map[string]int{1: {"1": 1, "2": 0}, 2: {"3": 3}, 3: nil}
		tests := []castTest[map[int]map[string]bool]{
			{nil, nil, ""},
			{map[string]map[bool]int{}, map[int]map[string]bool{}, ""},
			{map[int]map[bool]float64{}, map[int]map[string]bool{}, ""},
			{map[uint]map[int]string{1: {1: "true", 2: "false"}, 2: {3: "true"}, 3: nil}, expected, ""},
			{map[int16]map[uint]bool{1: {1: true, 2: false}, 2: {3: true}, 3: nil}, expected, ""},
			{map[uint64]map[int8]float64{1: {1: 1.5, 2: 0}, 2: {3: -2.3}, 3: nil}, expected, ""},
			{map[uint8]map[uint8]float32{1: {1: 1, 2: 0}, 2: {3: 3}, 3: nil}, expected, ""},
			{map[int32]map[int16]int8{1: {1: 1, 2: 0}, 2: {3: 3}, 3: nil}, expected, ""},
			{map[uint16]map[uint16]uint64{1: {1: 1, 2: 0}, 2: {3: 3}, 3: nil}, expected, ""},
			{map[float64]map[int32]int64{1: {1: 1, 2: 0}, 2: {3: 3}, 3: nil}, expected, ""},
			{map[int64]map[uint32]uint32{1: {1: 1, 2: 0}, 2: {3: 3}, 3: nil}, expected, ""},
			{map[float32]map[int64]int32{1: {1: 1, 2: 0}, 2: {3: 3}, 3: nil}, expected, ""},
			{map[int8]map[uint64]uint16{1: {1: 1, 2: 0}, 2: {3: 3}, 3: nil}, expected, ""},
			{map[uint8]map[float32]uint8{1: {1: 1, 2: 0}, 2: {3: 3}, 3: nil}, floatExpected, ""},
			{map[string]map[float64]int8{"1": {1: 1, 2: 0}, "2": {3: 3}, "3": nil}, floatExpected, ""},
			{map[bool]map[bool]uint{false: {true: 0}, true: {false: 1}}, map[int]map[string]bool{0: {"true": false}, 1: {"false": true}}, ""},
			{v0, expected, ""},
			{&v0, expected, ""},
			{map[string]string{"1": "invalid"}, nil, "failed to cast string to map[string]bool"},
			{"invalid", nil, "failed to cast string to map[int]map[string]bool"},
		}
		runCastTests(t, "AsMap[int, map[string]bool]", AsMap[int, map[string]bool], tests)
	})
}
