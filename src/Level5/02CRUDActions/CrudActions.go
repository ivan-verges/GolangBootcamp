package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID    string  `json:"Id"`
	Code  string  `json:"Code"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

type ProductInventory struct {
	Product  Product `json:"Product"`
	Quantity int     `json:"Available"`
}

func Add(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var productInventory ProductInventory
	json.Unmarshal(body, &productInventory)
	Inventory = append(Inventory, productInventory)

	json.NewEncoder(w).Encode(productInventory)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]
	for index, inventory := range Inventory {
		if inventory.Product.ID == id {
			Inventory = append(Inventory[:index], Inventory[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(Inventory)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Inventory)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]
	for _, inventory := range Inventory {
		if inventory.Product.ID == id {
			json.NewEncoder(w).Encode(inventory)
		}
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	body, _ := ioutil.ReadAll(r.Body)
	var productInventory ProductInventory
	json.Unmarshal(body, &productInventory)

	for index, inventory := range Inventory {
		if inventory.Product.ID == id {
			Inventory[index].Product = productInventory.Product
			Inventory[index].Quantity = productInventory.Quantity
			json.NewEncoder(w).Encode(Inventory[index])
		}
	}
}

// Global Variable to be Accessed by REST Methods
var Inventory []ProductInventory

func main() {
	Inventory = []ProductInventory{
		{Product: Product{ID: "1", Code: "0001", Name: "Product 1", Price: 12.15}, Quantity: 10},
		{Product: Product{ID: "2", Code: "0002", Name: "Product 2", Price: 10.95}, Quantity: 5},
		{Product: Product{ID: "3", Code: "0003", Name: "Product 3", Price: 11.25}, Quantity: 8},
		{Product: Product{ID: "4", Code: "0004", Name: "Product 4", Price: 15.30}, Quantity: 7},
		{Product: Product{ID: "5", Code: "0005", Name: "Product 5", Price: 9.29}, Quantity: 9},
	}

	fmt.Println("Serving Endpoints on: http://localhost:8888")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Inventory", Add).Methods(http.MethodPost)
	router.HandleFunc("/Inventory/{Id}", Delete).Methods(http.MethodDelete)
	router.HandleFunc("/Inventory", GetAll).Methods(http.MethodGet)
	router.HandleFunc("/Inventory/{Id}", GetById).Methods(http.MethodGet)
	router.HandleFunc("/Inventory/{Id}", Update).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8888", router))
}
