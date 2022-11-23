package quiz

import (
	"testing"
)

type Test struct {
	Min    int
	Max    int
	Values []int
	Want   []int
}

func TestFilterAge(t *testing.T) {
	test1 := Test{Min: 18, Max: 30, Values: []int{15, 18, 25, 30, 35}, Want: []int{18, 25, 30}}
	test2 := Test{Min: 20, Max: 40, Values: []int{15, 18, 25, 40, 55}, Want: []int{25, 40}}
	test3 := Test{Min: 75, Max: 100, Values: []int{-300, -100, 0, 74, 101, 500, 1000}, Want: []int{}} // Valores Fuera del Rango y Negativos
	test4 := Test{Min: -100, Max: -50, Values: []int{-15, -55, -75, 0, 25}, Want: []int{-55, -75}}    // Valores Negativo

	tests := []Test{test1, test2, test3, test4}

	for _, test := range tests {
		result := FilterAge(test.Min, test.Max, test.Values)
		equals := Equal(test.Want, result)
		if equals {
			t.Logf("TEST PASSED")
		} else {
			t.Errorf("TEST FAILED")
		}
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
