package storage

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"path/filepath"
	"sync"

	errors2 "github.com/pkg/errors"
	"github.com/twinj/uuid"
)

var (
	storagePath = flag.String("storagePath", "storage/Books.json", "path to the storage file")
	fileMutex   sync.Mutex
)

var (
	// For errors.
	errNoBookFound = errors.New("requested book doesn't exist")
)

func getPathToFS() (string, error) {
	path, err := filepath.Abs(*storagePath)
	if err != nil {
		return "", err
	}

	return path, nil

}

func writeData(books Books) error {
	// Will that work ?
	fileMutex.Lock()
	defer fileMutex.Unlock()
	path, err := getPathToFS()
	if err != nil {
		return err
	}

	booksBytes, err := json.MarshalIndent(books, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, booksBytes, 0644)
}

func readData() ([]byte, error) {
	path, err := getPathToFS()
	if err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

// GetBooks comment
func GetBooks() (Books, error) {
	var books Books

	raw, err := readData()
	if err != nil {
		return nil, errors2.Wrap(err, "Couldn't read data from storage")
	}

	return books, errors2.Wrap(json.Unmarshal(raw, &books), "Couldn't get books from storage")
}

// GetBookByID comment
func GetBookByID(id string) (Book, int, error) {

	var book Book
	books, err := GetBooks()

	if err != nil {
		return book, 0, errors2.Wrap(err, "Couldn't get book by ID")
	}

	for idx, book := range books {
		if book.ID == id {
			return book, idx, nil
		}
	}

	return book, 0, errNoBookFound
}

//AddBook comment
func AddBook(book Book) error {

	books, err := GetBooks()
	if err != nil {
		return err
	}

	book.ID = uuid.NewV4().String()
	books = append(books, book)

	return writeData(books)
}

//DeleteBook comment
func DeleteBook(id string) error {
	books, err := GetBooks()
	if err != nil {
		return err
	}

	_, idx, err := GetBookByID(id)

	if err != nil {
		return err
	}

	books = append(books[:idx], books[idx+1:]...)
	return writeData(books)

}

//UpdateBook comment
func UpdateBook(id string, updatedBook Book) error {

	books, err := GetBooks()
	if err != nil {
		return err
	}

	_, idx, err := GetBookByID(id)
	if err != nil {
		return err
	}

	book := &books[idx]
	book.Price = updatedBook.Price
	book.Title = updatedBook.Title
	book.Pages = updatedBook.Pages
	book.Genres = updatedBook.Genres

	return writeData(books)

}
