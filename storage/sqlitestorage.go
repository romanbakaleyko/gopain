package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
)

type sqliteStorage struct {
	storage *gorm.DB
}

func NewSqliteStorage(path string) (*sqliteStorage, error) {

	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't open database")
	}

	// Display SQL queries
	if err = db.LogMode(true).Error; err != nil {
		return nil, err
	}

	// Creating the table
	if !db.HasTable(&Book{}) {

		if err = db.CreateTable(&Book{}).Error; err != nil {
			return nil, errors.Wrap(err, "couldn't create table")
		}
	}

	return &sqliteStorage{storage: db}, nil
}

func (s *sqliteStorage) GetBooks() (Books, error) {
    var books Books
	return books, s.storage.Find(&books).Error
}

func (s *sqliteStorage) AddBook(book Book) error {
    book.InitBook()
	return s.storage.Create(&book).Error
}

//GetBookByID TODO: rework this, since for sqlite storage return index to update books doesn't make sense, need
// to create internal function for getting ID for filestorage.
func (s *sqliteStorage) GetBookByID(id string) (Book, int, error) {
	var book Book
	err := s.storage.Where("id = ?", id).First(&book).Error
	if err == gorm.ErrRecordNotFound {
		return book, 0, ErrNoBookFound
	}

	return book, 0, err
}

func (s *sqliteStorage) DeleteBook(id string) error {
	query := s.storage.Where("id = ?", id).Delete(&Book{})
	if query.Error != nil {
		return errors.Wrap(query.Error, "can't delete book")
	}

	if query.RowsAffected == 0 {
		return ErrNoBookFound
	}

	return nil
}

func (s *sqliteStorage) UpdateBook(id string, updatedBook Book) error {
	book, _, err := s.GetBookByID(id)
	if err != nil {
		return err
	}

	//TODO: rework this here and in the filestorage
	book.Price = updatedBook.Price
	book.Title = updatedBook.Title
	book.Pages = updatedBook.Pages
	book.Genres = updatedBook.Genres

	err = s.storage.Update(&book).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNoBookFound
	}

	return errors.Wrap(err, "couldn't update book")
}

