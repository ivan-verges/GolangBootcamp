package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Car struct {
	ID    string `json:"Id"`
	Make  string `json:"Make"`
	Model string `json:"Model"`
	Year  int    `json:"Year"`
}

type CarInventory struct {
	Car       Car `json:"Car"`
	Available int `json:"Available"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	carInventory := CarInventory{}
	json.Unmarshal(body, &carInventory)

	Inventory = append(Inventory, carInventory)

	json.NewEncoder(w).Encode(carInventory)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Inventory)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	for _, carInventory := range Inventory {
		if carInventory.Car.ID == id {
			json.NewEncoder(w).Encode(carInventory)
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	for index, carInventory := range Inventory {
		if carInventory.Car.ID == id {
			Inventory = append(Inventory[:index], Inventory[index+1:]...)
		}
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	body, _ := ioutil.ReadAll(r.Body)
	tmp := CarInventory{}
	json.Unmarshal(body, &tmp)

	for index, carInventory := range Inventory {
		if carInventory.Car.ID == id {
			Inventory[index].Car = tmp.Car
			Inventory[index].Available = tmp.Available

			json.NewEncoder(w).Encode(Inventory[index])
		}
	}

}

var Inventory []CarInventory

func main() {
	Inventory = []CarInventory{
		{Car: Car{ID: "1", Make: "Mitsubishi", Model: "Outlander", Year: 2015}, Available: 2},
		{Car: Car{ID: "2", Make: "Toyota", Model: "Corolla", Year: 2009}, Available: 4},
		{Car: Car{ID: "3", Make: "Honda", Model: "Civic", Year: 2011}, Available: 3},
	}

	fmt.Println("Serving on: http://localhost:8888")
	router := mux.NewRouter()
	router.HandleFunc("/Car", Create).Methods(http.MethodPost)
	router.HandleFunc("/Car", GetAll).Methods(http.MethodGet)
	router.HandleFunc("/Car/{Id}", GetById).Methods(http.MethodGet)
	router.HandleFunc("/Car/{Id}", Delete).Methods(http.MethodDelete)
	router.HandleFunc("/Car/{Id}", Update).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8888", router))
}
