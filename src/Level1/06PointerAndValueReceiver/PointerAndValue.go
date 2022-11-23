package main

import "fmt"

type MySlice1 struct {
	Indexes []int
}

func (slice MySlice1) AddIndex1(index int) {
	slice.Indexes = append(slice.Indexes, index)
}

func CreateMySlice1() MySlice1 {
	return MySlice1{Indexes: []int{}}
}

type MySlice2 struct {
	Indexes []int
}

func (slice *MySlice2) AddIndex2(index int) {
	slice.Indexes = append(slice.Indexes, index)
}

func CreateMySlice2() MySlice2 {
	return MySlice2{Indexes: []int{}}
}

func main() {
	slice1 := CreateMySlice1()
	slice1.AddIndex1(0)
	slice1.AddIndex1(1)
	fmt.Println(slice1.Indexes)

	slice2 := CreateMySlice2()
	slice2.AddIndex2(2)
	slice2.AddIndex2(3)
	fmt.Println(slice2.Indexes)
}
