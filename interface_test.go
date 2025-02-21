package cast

import "testing"

type eater interface {
	Eat()
}

type animal struct{}

func (a *animal) Eat() {}

func TestAsInterface(t *testing.T) {
	t.Run("as any", func(t *testing.T) {
		tests := []castTest[any]{
			{nil, nil, ""},
			{123, 123, ""},
			{"abc", "abc", ""},
			{true, true, ""},
			{[]int{1, 2, 3}, []int{1, 2, 3}, ""},
			{map[string]any{"a": 1, "b": false}, map[string]any{"a": 1, "b": false}, ""},
			{struct{}{}, struct{}{}, ""},
		}
		runCastTests(t, "AsInterface[any]", AsInterface[any], tests)
	})

	t.Run("as concrete interface", func(t *testing.T) {
		a := animal{}
		tests := []castTest[eater]{
			{nil, nil, ""},
			{&a, eater(&a), ""},
			{a, nil, "failed to cast cast.animal to cast.eater"},
		}
		runCastTests(t, "AsInterface[eater]", AsInterface[eater], tests)
	})
}
