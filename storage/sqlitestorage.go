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
	panic("not implemented")
	return nil, nil
}

func (s *sqliteStorage) AddBook(book Book) error {
	panic("not implemented")
	return nil
}

func (s *sqliteStorage) GetBookByID(id string) (Book, int, error) {
	panic("not implemented")
	return Book{}, 0, nil
}

func (s *sqliteStorage) DeleteBook(id string) error {
	panic("not implemented")
	return nil
}

func (s *sqliteStorage) UpdateBook(id string, updatedBook Book) error {
	panic("not implemented")
	return nil
}
