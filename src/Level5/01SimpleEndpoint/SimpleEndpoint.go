package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{'Message': 'Get Called'}"))
}

func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("{'Message': 'Post Called'}"))
}

func Generic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		{
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{'Message': 'Get Called'}"))
			break
		}
	case "POST":
		{
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("{'Message': 'Post Called'}"))
			break
		}
	default:
		{
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("{'Message': 'Unhandled Method Called'}"))
			break
		}
	}
}

func main() {
	fmt.Println("Serving Endpoints on: http://localhost:8888")
	//http.HandleFunc("/", Generic)
	//log.Fatal(http.ListenAndServe(":8888", nil))

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Get).Methods(http.MethodGet)
	router.HandleFunc("/", Post).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8888", router))
}
