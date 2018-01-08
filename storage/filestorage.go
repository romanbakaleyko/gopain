package storage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	errors2 "github.com/pkg/errors"
	"github.com/twinj/uuid"
)

var (
	fileMutex sync.Mutex
)

var (
	// For errors.
	errNoBookFound = errors.New("requested book doesn't exist")
)

type fileStorage struct {
	storage string
}

func NewFileStorage(path string) (*fileStorage, error) {

	storagePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	if _, err = os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(storagePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		_, err = file.WriteString("[]")
		if err != nil {
			return nil, err
		}

	}

	return &fileStorage{
		storage: storagePath,
	}, nil
}

func (s *fileStorage) writeData(books Books) error {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	booksBytes, err := json.MarshalIndent(books, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.storage, booksBytes, 0644)
}

func (s *fileStorage) readData() ([]byte, error) {

	raw, err := ioutil.ReadFile(s.storage)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

// GetBooks comment
func (s *fileStorage) GetBooks() (Books, error) {
	var books Books

	raw, err := s.readData()
	if err != nil {
		return nil, errors2.Wrap(err, "Couldn't read data from storage")
	}

	return books, errors2.Wrap(json.Unmarshal(raw, &books), "Couldn't get books from storage")
}

// GetBookByID comment
func (s *fileStorage) GetBookByID(id string) (Book, int, error) {

	var book Book
	books, err := s.GetBooks()

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
func (s *fileStorage) AddBook(book Book) error {

	books, err := s.GetBooks()
	if err != nil {
		return err
	}

	book.ID = uuid.NewV4().String()
	books = append(books, book)

	return s.writeData(books)
}

//DeleteBook comment
func (s *fileStorage) DeleteBook(id string) error {
	books, err := s.GetBooks()
	if err != nil {
		return err
	}

	_, idx, err := s.GetBookByID(id)

	if err != nil {
		return err
	}

	books = append(books[:idx], books[idx+1:]...)
	return s.writeData(books)

}

//UpdateBook comment
func (s *fileStorage) UpdateBook(id string, updatedBook Book) error {

	books, err := s.GetBooks()
	if err != nil {
		return err
	}

	_, idx, err := s.GetBookByID(id)
	if err != nil {
		return err
	}

	book := &books[idx]
	book.Price = updatedBook.Price
	book.Title = updatedBook.Title
	book.Pages = updatedBook.Pages
	book.Genres = updatedBook.Genres

	return s.writeData(books)

}
