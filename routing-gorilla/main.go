package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/book/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "you requested the book: %s on page %s\n", title, page)
	})

	// All HTTP verbs
	r.HandleFunc("/book/{title}", CreateBook).Methods("POST")
	r.HandleFunc("/book/{title}", ReadBook).Methods("GET").Schemes("https")
	r.HandleFunc("/book/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{title}", DeleteBook).Methods("DELETE")

	// hostname

	http.ListenAndServe(":5001", r)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Creting book %s", mux.Vars(r)["title"])
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Reading book %s", mux.Vars(r)["title"])
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Updating book %s", mux.Vars(r)["title"])
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Deleting book %s", mux.Vars(r)["title"])
}
