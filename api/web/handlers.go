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

type handler struct {
	storage Storage
}

type Storage interface {
	GetBooks() (storage.Books, error)
	GetBookByID(id string) (storage.Book, int, error)
	AddBook(book storage.Book) error
	DeleteBook(id string) error
	UpdateBook(id string, updatedBook storage.Book) error
}

func NewHandler(storage Storage) *handler {
	return &handler{
		storage: storage,
	}
}

// RootHandler comment
func (h *handler) RootHandler(w http.ResponseWriter, _ *http.Request) {

	_, err := fmt.Fprint(w, "Welcome to the library, to get more info use /helper URL")
	log.Info("Welcome to lib")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println(err)
	}
}

//HelperHandler
func (h *handler) HelperHandler(w http.ResponseWriter, r *http.Request) {

	_, err := fmt.Fprintf(w, helperMessage)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Println(err)
	}
}

// GetBooks handles GET
func (h *handler) GetBooksHandler(w http.ResponseWriter, r *http.Request) {

	books, err := h.storage.GetBooks()

	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *handler) validateID(r *http.Request) (string, error) {

	vars := mux.Vars(r)
	id, ok := vars["id"]

	_, err := uuid.Parse(id)

	if !ok || err != nil {
		return "", err
	}

	return id, nil

}

func (h *handler) validateBookFields(r *http.Request) (storage.Book, error) {

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
func (h *handler) GetBookHandler(w http.ResponseWriter, r *http.Request) {

	id, err := h.validateID(r)

	if err != nil {
		log.Info("Bad request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	book, _, err := h.storage.GetBookByID(id)
	if err != nil {
		if err == storage.ErrNoBookFound {
			log.Info(err)
			w.WriteHeader(http.StatusNotFound)
		}
		log.Warn(err)
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
func (h *handler) AddBookHandler(w http.ResponseWriter, r *http.Request) {

	book, err := h.validateBookFields(r)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = h.storage.AddBook(book)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

//DeleteBook
func (h *handler) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {

	id, err := h.validateID(r)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.storage.DeleteBook(id)
	if err != nil {
		if err == storage.ErrNoBookFound {
			log.Info(err)
			w.WriteHeader(http.StatusNotFound)
		}
		log.Warn(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

//GetFilteredBooksHandler
func (h *handler) GetFilteredBooksHandler(w http.ResponseWriter, r *http.Request) {}

//UpdateBookHandler
func (h *handler) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {

	id, err := h.validateID(r)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// This is not what I wanted to get. Need to discuss.
	book, err := h.validateBookFields(r)
	if err != nil {
		log.Info(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	err = h.storage.UpdateBook(id, book)
	if err != nil {
		if err == storage.ErrNoBookFound {
			log.Info(err)
			w.WriteHeader(http.StatusNotFound)
		}
		log.Warn(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
