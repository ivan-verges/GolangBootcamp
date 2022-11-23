package main

type Product struct {
	Name string
	Cost float64
}

func (product *Product) UpdateCost(value float64) {
	product.Cost = value
}

func main() {

}
