package Multiple

import (
	"fmt"
	"testing"
)

type Test struct {
	A    int
	B    int
	Want int
}

func TestIntMin(t *testing.T) {
	tests := []Test{}
	tests = append(tests, Test{A: 0, B: 1, Want: 0})
	tests = append(tests, Test{A: 1, B: 0, Want: 0})
	tests = append(tests, Test{A: -5, B: -10, Want: -10})
	tests = append(tests, Test{A: -25, B: 25, Want: -25})
	tests = append(tests, Test{A: -1, B: 0, Want: 0}) //Failing Test

	for _, test := range tests {
		testname := fmt.Sprintf("IntMin(%d,%d)", test.A, test.B)
		t.Run(testname, func(t *testing.T) {
			result := IntMin(test.A, test.B)
			if result != test.Want {
				t.Errorf("IntMin(%d, %d) FAILED. WANT: %d, GOT: %d", test.A, test.B, test.Want, result)
			} else {
				t.Logf("IntMin(%d, %d) PASSED. WANT: %d, GOT: %d", test.A, test.B, test.Want, result)
			}
		})
	}
}
