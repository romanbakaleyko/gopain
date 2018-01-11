package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/romanbakaleyko/gopain/storage"
	log "github.com/sirupsen/logrus"
	"github.com/twinj/uuid"
)

const helperMessage = `This is short helper to help you query beter:
*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*==*
	
	[/] [GET]		Welcome page
	[/helper] [GET]		Helper
	[/books] [GET]		Get all books from the lib
	[/books] [POST]		Add a book; fields: {Title:string, Genres: []string, Pages: int, Price: float32}
	[/books/filter] [POST]  	Get all books from the lib and filter them
	[/books/[id]] [GET]   	Get a book by ID
	[/books/[id]] [PUT]   	Update existing book by ID
	[/books/[id]] [DELETE]	Delete a book by ID

`

var (
	errMissedGenre = errors.New("bad input, missing values for field genre")
	errMissedPages = errors.New("bad input, missing values for field pages")
	errMissedPrice = errors.New("bad input, missing values for field price")
	errMissedTitle = errors.New("bad input, missing values for field title")
)

// RootHandler comment
func RootHandler(w http.ResponseWriter, _ *http.Request) {
	logger := log.WithField("handler", "RootHandler")

	_, err := fmt.Fprint(w, "Welcome to the library, to get more info use /helper URL")
	logger.Info("Welcome to lib")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Warn(err)
	}
}

//HelperHandler
func HelperHandler(w http.ResponseWriter, r *http.Request) {

	_, err := fmt.Fprintf(w, helperMessage)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Warn(err)
	}
}

// GetBooks handles GET
func GetBooksHandler(w http.ResponseWriter, r *http.Request) {

	books, err := storage.GetBooks()

	if err != nil {
		// TODO: handle internal error
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func validateID(r *http.Request) (string, error) {

	vars := mux.Vars(r)
	id := vars["id"]

	_, err := uuid.Parse(id)
	return id, err
}

// TODO: move validation outside handler package
func validateBookFields(r *http.Request) (storage.Book, error) {

	var book storage.Book
	err := errors.Wrap(json.NewDecoder(r.Body).Decode(&book), "Couldn't decode body")
	if err != nil {
		log.Info(err)
		return book, err
	}

	switch {
	case book.Genres == nil:
		return book, errMissedGenre
	case book.Pages == 0:
		return book, errMissedPages
	case book.Price == 0:
		return book, errMissedPrice
	case book.Title == "":
		return book, errMissedTitle
	}

	return book, nil
}

// GetBook handles GET
func GetBookHandler(w http.ResponseWriter, r *http.Request) {

	id, err := validateID(r)

	if err != nil {
		// TODO: Use typed error from validateID
		log.Info("Bad request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, _, err := storage.GetBookByID(id)
	if err != nil {
		if err == storage.ErrNoBookFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//CreateBook
func AddBookHandler(w http.ResponseWriter, r *http.Request) {

	book, err := validateBookFields(r)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = storage.AddBook(book)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

//DeleteBook
func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {

	id, err := validateID(r)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = storage.DeleteBook(id)
	if err != nil {
		// TODO: handle "Not found"
		log.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

}

//GetFilteredBooksHandler
func GetFilteredBooksHandler(w http.ResponseWriter, r *http.Request) {}

//UpdateBookHandler
func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	// https://github.com/8tomat8/go-talks/blob/master/code-from-class/update-struct-from-json/main.go
	id, err := validateID(r)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// This is not what I wanted to get. Need to discuss.
	book, err := validateBookFields(r)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = storage.UpdateBook(id, book)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
