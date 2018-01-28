package storage

import (
	"github.com/lib/pq"
	"github.com/twinj/uuid"
)

//Book comment
type Book struct {
	ID     string         `gorm:"type:varchar(100);primary_key" json:"id, omitempty"`
	Title  string         `gorm:"type:varchar(100)" json:"title, omitempty"`
	Genres pq.StringArray `gorm:"type:varchar(64)" json:"genres, omitempty"`
	Pages  int            `gorm:"type:int" json:"pages, omitempty"`
	Price  float32        `gorm:"type:real" json:"price, omitempty"`
}

//Books comment
type Books []Book

//BookFilter comment
type BookFilter struct {
	Price string `json:"price, omitempty"`
}

//Validate Book fields
func (b *Book) ValidateBookFields() error {

	if b.Genres == nil {
		return ErrMissedGenre
	}
	if  b.Pages == 0 {
		return ErrMissedPages
	}
	if b.Price == 0 {
		return ErrMissedPrice
	}
	if b.Title == "" {
		return ErrMissedTitle
	}
	return nil
}

//InitBook Book
func (b *Book) InitBook() {
	b.ID = uuid.NewV4().String()
}
//
//func (b Book) ValidateBookId() error{
//
//	vars := mux.Vars(r)
//	id := vars["id"]
//
//	_, err := uuid.Parse(id)
//
//	return id, err
//}
