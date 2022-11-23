package main

import (
	"fmt"
)

type Product struct {
	Id   string
	Name string
}

type Inventory struct {
	Data map[string]Product
}

func NewInventory() Inventory {
	return Inventory{make(map[string]Product)}
}

func NewProduct(Id string, Name string) Product {
	return Product{Id: Id, Name: Name}
}

func (product Product) ShowProduct() {
	fmt.Println("Id:", product.Id)
	fmt.Println("Name:", product.Name)
}

func (inventory Inventory) ShowInventory() {
	for key, value := range inventory.Data {
		fmt.Println("Key:", key)
		value.ShowProduct()
	}
}

func main() {
	//product0 := NewProduct("", "Bolt")
	product1 := NewProduct("1", "Bolt")
	product2 := NewProduct("2", "Drill")
	product3 := NewProduct("3", "Hammer")
	product4 := NewProduct("4", "Saw")
	product5 := NewProduct("5", "Screwdriver")

	inventory := NewInventory()
	//inventory = AddProductToInventory(&inventory, product0) //Empty ID to Raise a Panic Alert
	inventory = AddProductToInventory(&inventory, product1)
	inventory = AddProductToInventory(&inventory, product2)
	inventory = AddProductToInventory(&inventory, product3)
	inventory = AddProductToInventory(&inventory, product4)
	inventory = AddProductToInventory(&inventory, product5)
	//inventory = AddProductToInventory(&inventory, product5) //Duplicated ID to Raise a Panic Alert

	inventory.ShowInventory()
}

// Add Error Return to Handle Exception (Panic)
func AddProductToInventory(inventory *Inventory, product Product) Inventory {

	if len(product.Id) <= 0 {
		panic("Can't Register a Product With an Empty ID")
	} else if _, ok := inventory.Data[product.Id]; ok {
		panic("Product Id is Duplicated")
	} else {
		inventory.Data[product.Id] = product
	}

	return *inventory
}
