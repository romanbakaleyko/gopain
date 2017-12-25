package web

import (
	"log"
	"net/http"
)

// GetBooks
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Get books")
}

// GetBook
func GetBookHandler(w http.ResponseWriter, r *http.Request) {}

//CreateBook
func AddBookHandler(w http.ResponseWriter, r *http.Request) {

}

//DeleteBook
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {}
