package Single

import (
	"testing"
)

func TestIntMin(t *testing.T) {
	result := IntMin(7, 5)
	if result != 5 {
		t.Errorf("IntMin(5, 7) FAILED. WANT: 5, GOT: %d", result)
	} else {
		t.Logf("IntMin(5, 7) PASSED. WANT: 5, GOT: %d", result)
	}
}
