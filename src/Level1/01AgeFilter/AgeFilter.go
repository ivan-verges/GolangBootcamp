package main

import (
	"fmt"
)

func main() {
	min := 18
	max := 65
	data := []int{10, 25, 32, 14, 18, 63, 78, 65, 92, 54}

	result := FilterAge(min, max, data)

	fmt.Println(result)
}

func FilterAge(min int, max int, data []int) []int {

	result := []int{}

	for _, value := range data {
		if IsBetweenRange(value, min, max) {
			result = append(result, value)
		}
	}

	return result
}

func IsBetweenRange(number int, min int, max int) bool {
	if number >= min && number <= max {
		return true
	} else {
		return false
	}
}
