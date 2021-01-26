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

	//Agregando base de datos interna
	books = append(books, Book{ID: "1", Isbn: "123456", Title: "It", Author: &Author{Firstname: "Stephen", Lastname: "King"}})
	books = append(books, Book{ID: "2", Isbn: "789456", Title: "Star wars ", Author: &Author{Firstname: "George ", Lastname: "Lucas"}})
	books = append(books, Book{ID: "3", Isbn: "987456", Title: "Star Trek", Author: &Author{Firstname: "Gene", Lastname: "Roddenberry"}})
	books = append(books, Book{ID: "4", Isbn: "321456", Title: "Alien: el octavo pasajero", Author: &Author{Firstname: "Ridley", Lastname: "Scott"}})

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
	Firstname string `json:"Firstname"` //nombre
	Lastname  string `json:"Lastname"`  //apellido
}

// Libros Estructura (modelos)
type Book struct {
	ID     string  `json:"id"`     //id identificador
	Isbn   string  `json:"isbn"`   //Identificador unico para libros
	Title  string  `json:"title"`  //titulo
	Author *Author `json:"author"` //autor
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

// Funcion para eliminar Libros
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Parametros Mux
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Request Simple
// {
//	"id" : "5",
// 	"isbn único para libros":"123456",
// 	"title":"El señor de los Anillos",
// 	"author":{"firstname":"J. R. R",
//            "lastname":"Tolkien"
//			 }
// }
