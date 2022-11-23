package main

import (
	"fmt"
	"math"
)

func main() {
	added := AddNumbers(5, 4)
	substracted := SubstractNumbers(6, 2)
	multiplied := MultiplyNumbers(5, 9)
	divided := DivideNumbers(4, 2)
	powered := PowerNumber(2, 3)
	squared := SquareRootNumber(25)

	fmt.Println(added)
	fmt.Println(substracted)
	fmt.Println(multiplied)
	fmt.Println(divided)
	fmt.Println(powered)
	fmt.Println(squared)
}

func AddNumbers(number1 float64, number2 float64) float64 {
	return number1 + number2
}

func SubstractNumbers(number1 float64, number2 float64) float64 {
	return number1 - number2
}

func MultiplyNumbers(number1 float64, number2 float64) float64 {
	return number1 * number2
}

func DivideNumbers(number1 float64, number2 float64) float64 {

	if number2 == 0 {
		panic("Can't Divide by Zero")
	} else {
		return number1 / number2
	}
}

func PowerNumber(number1 float64, number2 float64) float64 {
	return math.Pow(number1, number2)
}

func SquareRootNumber(number1 float64) float64 {
	if number1 < 0 {
		panic("Can't Calculate Square Root of Negative Numbers")
	} else {
		return math.Sqrt(number1)
	}
}
