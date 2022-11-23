package main

import (
	"fmt"
)

func main() {
	number := 5

	ByValue(number)
	fmt.Println(number)

	ByReference(&number)
	fmt.Println(number)
}

func ByValue(input int) {
	input += 10
}

func ByReference(input *int) {
	*input += 10
}
