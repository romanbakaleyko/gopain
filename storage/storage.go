package storage

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/twinj/uuid"
)

//Book comment
type Book struct {
	ID     string   `json:"id, omitempty"`
	Title  string   `json:"title, omitempty"`
	Genres []string `json:"genre, omitempty"`
	Pages  int      `json:"pages, omitempty"`
	Price  float32  `json:"price, omitempty"`
}

//Books comment
type Books []Book

var storagePath = flag.String("storagePath", "storage/Books.json", "path to the storage file")

func getPathToFS() (string, error) {
	path, err := filepath.Abs(*storagePath)
	if err != nil {
		return "", err
	}

	return path, nil

}

func writeData(books Books) error {
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
	log.Println(">>> Getting books from a storage.")
	var books Books

	raw, err := readData()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(raw, &books); err != nil {
		return nil, err
	}

	return books, nil
}

// GetBookByID comment
func GetBookByID(id string) {}

//AddBook comment
func AddBook(book Book) error {
	err := errors.New("bad input, missing values for some fields")
	switch {
	case book.Genres == nil:
		return err
	case book.Pages == 0:
		return err
	case book.Price == 0:
		return err
	case book.Title == "":
		return err
	}

	books, err := GetBooks()
	if err != nil {
		return err
	}

	book.ID = uuid.NewV4().String()
	books = append(books, book)

	return writeData(books)
}

//DeleteBook comment
func DeleteBook() {}
