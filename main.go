package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

// Agregando Estructuras

// Autor Estructura
type Author struct {
	Firstname string `json:"nombre"`
	Lastname  string `json:"apellido"`
}

// Libros Estructura (modelos)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"identificador único para libros"`
	Title  string  `json:"titulo"`
	Author *Author `json:"autor"`
}

//Variable, Inicializar Books
var books []Book

//Funcion para Obtener todos los Libros
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Funcion para Obtener un Solo libro
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Recorre los libros y encuentra uno con la identificación de los parámetros
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Funcion para Agregar un libro
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//Funcion para Actualizar libros
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}
