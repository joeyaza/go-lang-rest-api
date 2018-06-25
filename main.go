package main

import(
"encoding/json"
"log"
"net/http"
"math/rand"
"strconv"
"github.com/gorilla/mux"
)

// struct - like a class

type Book struct {

	ID 		string	`json:"id"`
	Isbn 	string	`json:"isbn"`
	Title 	string	`json:"title"`
	Author	*Author	`json:"author"`

}

type Author struct {

	Firstname	string	`json:"firstname"`
	Lastname	string 	`json:"lastname"`

}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // gets params

	for _, item:= range books {

		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)

			return

		}

	}

	json.NewEncoder(w).Encode(&Book{})

	
}

func createBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000)) //mockid

	books = append(books, book)

	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // gets params

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

func deleteBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // gets params

	for index, item := range books {

		if item.ID == params["id"] {

		books = append(books[:index], books[index+1:]...)
		break

		}

	}

	json.NewEncoder(w).Encode(books)

	
}

func main() {

	//init router
	r := mux.NewRouter()

	//mockdata
	books = append(books, Book{ID: "1", Isbn: "mock123", Title: "Book one", Author: &Author{
			Firstname: "Joe", Lastname: "Az"}})
	books = append(books, Book{ID: "2", Isbn: "mock456", Title: "Book two", Author: &Author{
			Firstname: "Jon", Lastname: "Noaz"}})

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}