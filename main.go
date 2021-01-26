package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Iniciando Routers
	r := mux.NewRouter()

	// Handlers y endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Iniciando Servidor
	log.Fatal(http.ListenAndServe(":8000", r))
}
