package web

import (
	"github.com/gopain/storage"
	"github.com/gorilla/mux"
)

// CreateRoutes comment
func CreateRoutes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/books", storage.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", storage.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", storage.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", storage.DeleteBook).Methods("DELETE")

	return router

}
