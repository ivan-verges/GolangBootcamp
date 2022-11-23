package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (shape Circle) Area() float64 {
	return math.Pi * shape.Radius * shape.Radius
}

func (shape Circle) Perimeter() float64 {
	return 2 * (math.Pi * shape.Radius)
}

func CreateCircle(radius float64) Shape {
	return Circle{Radius: radius}
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (shape Rectangle) Area() float64 {
	return shape.Height * shape.Width
}

func (shape Rectangle) Perimeter() float64 {
	return 2 * (shape.Height + shape.Width)
}

func CreateRectangle(height float64, width float64) Shape {
	return Rectangle{Height: height, Width: width}
}

func main() {
	circle := CreateCircle(3)
	rectangle := CreateRectangle(4, 5)

	fmt.Println("Circle Details:")
	PrintShapeArea(circle)
	PrintShapePerimeter(circle)

	fmt.Println("Rectangle Details:")
	PrintShapeArea(rectangle)
	PrintShapePerimeter(rectangle)
}

func PrintShapeArea(shape Shape) {
	fmt.Println("Area:", shape.Area())
}

func PrintShapePerimeter(shape Shape) {
	fmt.Println("Perimeter:", shape.Perimeter())
}
