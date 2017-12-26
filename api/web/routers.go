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
	router.HandleFunc("/book/{id}", GetBookHandler).Methods("GET")
	router.HandleFunc("/book/{id}", AddBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", DeleteBookHandler).Methods("DELETE")
	router.HandleFunc("/book/{id}", UpdateBookHandler).Methods("PUT")
	router.HandleFunc("/books/filter", GetFilteredBooksHandler).Methods("POST")

	return router
}
