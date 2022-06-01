package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new book")
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Read a book")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update a book")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete a book")
}

func SecureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Secure endpoint")
}

func InsecureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Insecure endpoint")
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Book handler")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/book/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	r.HandleFunc("/books/{title}", CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/books/{title}", ReadBook).Methods(http.MethodGet)
	r.HandleFunc("/books/{title}", UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books/{title}", DeleteBook).Methods(http.MethodDelete)

	r.HandleFunc("/books/{title}", BookHandler).Host("www.amirsorouri00.com")

	r.HandleFunc("/secure", SecureHandler).Schemes("https")
	r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	http.ListenAndServe(":6060", r)
}
