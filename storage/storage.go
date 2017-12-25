package storage

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Book struct {
	id     string   `json:"id, omitempty"`
	title  string   `json:"title, omitempty"`
	ganres []string `json:"genre, omitempty"`
	pages  int      `json:"pages, omitempty"`
	price  float32  `json:"price, omitempty"`
}

type Books []Book

var storagePath = flag.String("storagePath", "Books.json", "path to the storage file")

func ReadFileStorage() []byte {
	path, err := filepath.Abs(*storagePath)

	if err != nil {
		log.Println(err)
	}

	raw, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println(err)
	}

	return raw
}

// GetBooks
func GetBooks() {

	var books Books
	raw := ReadFileStorage()

	err := json.Unmarshal(raw, &books)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(books)

}

// GetBook
func GetBookById(id string) Book {}

//CreateBook
func AddBook() {}

//DeleteBook
func DeleteBook() {}

func main() {
	GetBooks()

}
