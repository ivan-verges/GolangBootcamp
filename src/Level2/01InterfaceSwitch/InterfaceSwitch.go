package main

import (
	"fmt"
	"reflect"
)

type BaseType interface {
	GetTypeName() string
}

type MyString struct {
	Value string
}

func (input MyString) GetTypeName() string {
	return reflect.ValueOf(input.Value).Kind().String()
}

type MyInt struct {
	Value int
}

func (input MyInt) GetTypeName() string {
	return reflect.ValueOf(input.Value).Kind().String()
}

type MyBool struct {
	Value bool
}

func (input MyBool) GetTypeName() string {
	return reflect.ValueOf(input.Value).Kind().String()
}

func main() {
	ms := MyString{Value: "Hello"}
	mi := MyInt{Value: 25}
	mb := MyBool{Value: true}

	PrintType(ms)
	PrintType(mi)
	PrintType(mb)
}

func PrintType(input BaseType) {
	/*
		ib, ob := input.(MyBool)
		if ob {
			fmt.Println("Type: Bool, Value:", ib)
		}

		ii, oi := input.(MyInt)
		if oi {
			fmt.Println("Type: Int, Value:", ii)
		}

		is, os := input.(MyString)
		if os {
			fmt.Println("Type: String, Value:", is)
		}
	*/

	//t := input.GetTypeName()
	t := reflect.TypeOf(input).String()
	//input.(type)
	switch t {
	case "bool":
		fmt.Println("Bool")
		break
	case "main.MyBool":
		fmt.Println("MyBool => Bool")
		break
	case "int":
		fmt.Println("Int")
		break
	case "main.MyInt":
		fmt.Println("MyInt => Int")
		break
	case "string":
		fmt.Println("String")
		break
	case "main.MyString":
		fmt.Println("MyString => String")
		break
	default:
		fmt.Println(t)
	}
}
