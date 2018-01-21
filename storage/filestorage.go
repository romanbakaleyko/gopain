package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
	"github.com/twinj/uuid"
)

type fileStorage struct {
	storage   string
	fileMutex sync.RWMutex
}

func NewFileStorage(path string) (*fileStorage, error) {

	storagePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	if _, err = os.Stat(path); err == nil {
		return &fileStorage{storage: storagePath}, nil
	}

	// If file not exist
	file, err := os.Create(storagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	_, err = file.WriteString("[]")
	if err != nil {
		return nil, err
	}

	return &fileStorage{storage: storagePath}, nil
}

func (s *fileStorage) writeData(books Books) error {
	booksBytes, err := json.MarshalIndent(books, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.storage, booksBytes, 0644)
}

func (s *fileStorage) readData() ([]byte, error) {
	return ioutil.ReadFile(s.storage)
}

// GetBooks comment
func (s *fileStorage) GetBooks() (Books, error) {
	s.fileMutex.RLock() // Write
	defer s.fileMutex.RUnlock()
	return s.getBooks()
}

func (s *fileStorage) getBooks() (Books, error) {
	var books Books

	raw, err := s.readData()
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't read data from storage")
	}

	return books, errors.Wrap(json.Unmarshal(raw, &books), "Couldn't get books from storage")
}

// GetBookByID comment
func (s *fileStorage) GetBookByID(id string) (Book, int, error) {
	var book Book
	books, err := s.GetBooks()

	if err != nil {
		return book, 0, errors.Wrap(err, "Couldn't get book by ID")
	}

	for idx, book := range books {
		if book.ID == id {
			return book, idx, nil
		}
	}

	return book, 0, ErrNoBookFound
}

//AddBook comment
func (s *fileStorage) AddBook(book Book) error {
	s.fileMutex.Lock() // Read/Write
	defer s.fileMutex.Unlock()

	books, err := s.getBooks()
	if err != nil {
		return err
	}

	book.ID = uuid.NewV4().String()
	books = append(books, book)

	return s.writeData(books)
}

//DeleteBook comment
func (s *fileStorage) DeleteBook(id string) error {
	s.fileMutex.Lock() // Read/Write
	defer s.fileMutex.Unlock()

	books, err := s.getBooks()
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
	s.fileMutex.Lock() // Read/Write
	defer s.fileMutex.Unlock()

	books, err := s.getBooks()
	if err != nil {
		return err
	}

	_, idx, err := s.GetBookByID(id)
	if err != nil {
		return err
	}

	// What would happen if user omits some field?
	book := &books[idx]
	book.Price = updatedBook.Price
	book.Title = updatedBook.Title
	book.Pages = updatedBook.Pages
	book.Genres = updatedBook.Genres

	return s.writeData(books)
}
