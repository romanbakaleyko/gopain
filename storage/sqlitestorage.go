package storage

import (
	"github.com/jinzhu/gorm"
)

type sqliteStorage struct {
	storage *gorm.DB
}

func NewSqliteStorage(path string) (*sqliteStorage, error) {

	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Display SQL queries
	if err = db.LogMode(true).Error; err != nil {
		return nil, err
	}

	// Creating the table
	if !db.HasTable(&Book{}) {

		if err = db.CreateTable(&Book{}).Error; err != nil {
			return nil, err
		}
		if err = db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Book{}).Error; err != nil {
			return nil, err
		}
	}

	return &sqliteStorage{
		storage: db,
	}, nil
}

func (s *sqliteStorage) GetBooks() (Books, error) {
	return nil, nil

}

func (s *sqliteStorage) AddBook(book Book) error {
	return nil
}

func (s *sqliteStorage) GetBookByID(id string) (Book, int, error) {
	return Book{}, 0, nil
}

func (s *sqliteStorage) DeleteBook(id string) error {
	return nil
}

func (s *sqliteStorage) UpdateBook(id string, updatedBook Book) error {
	return nil
}
