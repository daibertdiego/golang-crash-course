package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct - Model

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

// Init books var as a slice Book struct

var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	jsonData, err := json.Marshal(books)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(jsonData))
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(404)
	json.NewEncoder(w).Encode("Book Not Found")
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Invalid data for Book creation")
		return
	}
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update a single Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			var book Book
			err := json.NewDecoder(r.Body).Decode(&book)
			if err != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode("Invalid data for Book update")
				return
			}
			books = append(books[:index], books[index+1:]...)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Delete a single Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
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
	// Init Router
	router := mux.NewRouter()

	// Mock Data
	books = append(books,
		Book{
			ID:    "1",
			Isbn:  "131313",
			Title: "O Abusado",
			Author: &Author{
				FirstName: "Caco",
				LastName:  "Barcelos",
				Age:       57,
			},
		},
	)

	books = append(books,
		Book{
			ID:    "2",
			Isbn:  "13112",
			Title: "Clean Code",
			Author: &Author{
				FirstName: "Robert",
				LastName:  "Martin",
				Age:       66,
			},
		},
	)

	books = append(books,
		Book{
			ID:    "3",
			Isbn:  "2314433",
			Title: "Lord of the Rings",
			Author: &Author{
				FirstName: "Stephen",
				LastName:  "King",
				Age:       99,
			},
		},
	)

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Default().Println("App up and running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
