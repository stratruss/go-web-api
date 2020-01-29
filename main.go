package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

const (
	contentType     = "content-Type"
	applicationJSON = "application/json"
)

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)
	params := mux.Vars(r)

	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update a Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)

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
	json.NewEncoder(w).Encode(books)
}

// Delete a Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// ルーターの初期化
	r := mux.NewRouter()

	// モックデータの作成
	books = append(books, Book{ID: "1", Title: "Book one", Author: &Author{FirstName: "Philip", LastName: "Williams"}})
	books = append(books, Book{ID: "2", Title: "Book Two", Author: &Author{FirstName: "John", LastName: "Johnson"}})

	// ルート（エンドポイント）
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
