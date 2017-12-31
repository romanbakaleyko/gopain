package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/romanbakaleyko/gopain/storage"
)

// RootHandler
func RootHandler(w http.ResponseWriter, _ *http.Request) {

	_, err := fmt.Fprint(w, "Welcome to the library, to get more info use /helper URL")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println(err)
	}
}

//
func HelperHandler(w http.ResponseWriter, r *http.Request) {

	msg := `This is short helper to help you query beter:
*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*
	
	[/] [GET]		Welcome page
	[/helper] [GET]		Helper
	[/books] [GET]		Get all books from the lib
	[/books/filter] [POST]  	Get all books from the lib and filter them
	[/book/[id]] [GET]   	Get a book by ID
	[/book/[id]] [POST]		Add a book
	[/book/[id]] [PUT]   	Update existing book by ID
	[/book/[id]] [DELETE]	Delete a book by ID


`

	_, err := fmt.Fprintf(w, msg)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println(err)
	}
}

// GetBooks
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	storage.GetBooks()
}

// GetBook
func GetBookHandler(w http.ResponseWriter, r *http.Request) {}

//CreateBook
func AddBookHandler(w http.ResponseWriter, r *http.Request) {}

//DeleteBook
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {}

//GetFilteredBooksHandler
func GetFilteredBooksHandler(w http.ResponseWriter, r *http.Request) {}

//UpdateBookHandler
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {}
