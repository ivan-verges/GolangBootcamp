package main

import (
	"fmt"
)

func main() {
	Data := []int{1, 2, 3, 4, 5}

	Data = Enqueue(Data, 6)
	fmt.Println(Data)

	Data = Dequeue(Data, 5)
	Data = Dequeue(Data, 3)
	Data = Dequeue(Data, 1)
	fmt.Println(Data)
}

func Enqueue(data []int, input int) []int {
	data = append(data, input)
	return data
}

func Dequeue(data []int, index int) []int {
	return append(data[:index], data[index+1:]...)
}
