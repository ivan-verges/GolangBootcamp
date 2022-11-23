package Different

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

type Test struct {
	Flag bool
	Want string
}

func TestStringOk(t *testing.T) {
	tests := []Test{}
	tests = append(tests, Test{Flag: true, Want: "Ok"})
	tests = append(tests, Test{Flag: false, Want: ""})   //Failing Assert (Expected Err to be Nil)
	tests = append(tests, Test{Flag: true, Want: "OK"})  //Failing Assert (Expected "OK" but got "Ok")
	tests = append(tests, Test{Flag: false, Want: "OK"}) //Failing Assert (Expected Err to be Nil, and "OK" but got "Ok")

	for index, test := range tests {
		result, err := StringOk(test.Flag)

		fmt.Println("Test:", index)
		assert.Equal(t, nil, err, "Expected Error to be Nil")
		assert.NotEqual(t, "Empty String", err, "Expected Error to be Nil")
		assert.Equal(t, test.Want, result, "Result Value is Different from Expected")
	}
}
