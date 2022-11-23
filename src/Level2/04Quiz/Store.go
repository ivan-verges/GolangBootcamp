package main

/*
A store sells two types of products: books and games. Each product has the fields name (string) and price (float). Define the following functionalities:
1. Product has the following methods:
   - A method to print the information of each product (type, name, and price)
   - A method to apply a discount ratio to the price
2. The store should be able to apply custom discounts based on the type of product: 10% discount for books and 20% discount for games.
*/

import (
	"fmt"
)

type Product interface {
	Print()
	Discount() float64
}

type Book struct {
	Name  string
	Price float64
}

func (product Book) Print() {
	fmt.Println("Book:", product.Name, "Price:", product.Price)
}

func (product Book) Discount() float64 {
	return product.Price * 0.9
}

func CreateBook(name string, price float64) Product {
	return Book{Name: name, Price: price}
}

type Game struct {
	Name  string
	Price float64
}

func (product Game) Print() {
	fmt.Println("Game:", product.Name, "Price:", product.Price)
}

func (product Game) Discount() float64 {
	return product.Price * 0.8
}

func CreateGame(name string, price float64) Product {
	return Game{Name: name, Price: price}
}

func main() {
	book1 := CreateBook("Libro 1", 2.5)
	book2 := CreateBook("Libro 2", 4.8)
	book3 := CreateBook("Libro 3", 9.2)
	game1 := CreateGame("Game 1", 7.9)
	game2 := CreateGame("Game 2", 8.1)
	game3 := CreateGame("Game 3", 9.75)

	products := []Product{}
	products = append(products, book1)
	products = append(products, book2)
	products = append(products, book3)
	products = append(products, game1)
	products = append(products, game2)
	products = append(products, game3)

	for _, product := range products {
		product.Print()
		fmt.Println("Con Descuento:", product.Discount())
	}
}
