package web

import (
	"github.com/gorilla/mux"
)

// CreateRoutes comment
func CreateRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", RootHandler).Methods("GET")
	router.HandleFunc("/helper", HelperHandler).Methods("GET")
	router.HandleFunc("/books", GetBooksHandler).Methods("GET")
	router.HandleFunc("/books", AddBookHandler).Methods("POST")
	router.HandleFunc("/books/{id}", GetBookHandler).Methods("GET")
	router.HandleFunc("/books/{id}", DeleteBookHandler).Methods("DELETE")
	router.HandleFunc("/books/{id}", UpdateBookHandler).Methods("PUT")
	router.HandleFunc("/books/filter", GetFilteredBooksHandler).Methods("POST")

	return router
}
