package web

import (
	"github.com/gorilla/mux"
)

// CreateRoutes comment
func CreateRoutes(handler *handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.RootHandler).Methods("GET")
	router.HandleFunc("/helper", handler.HelperHandler).Methods("GET")
	router.HandleFunc("/books", handler.GetBooksHandler).Methods("GET")
	router.HandleFunc("/books", handler.AddBookHandler).Methods("POST")
	router.HandleFunc("/books/{id}", handler.GetBookHandler).Methods("GET")
	router.HandleFunc("/books/{id}", handler.DeleteBookHandler).Methods("DELETE")
	router.HandleFunc("/books/{id}", handler.UpdateBookHandler).Methods("PUT")
	router.HandleFunc("/books/filter", handler.GetFilteredBooksHandler).Methods("POST")

	return router
}
