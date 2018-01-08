package storage

import "github.com/lib/pq"

//Book comment
type Book struct {
	ID     string         `gorm:"type:varchar(100);primary_key" json:"id, omitempty"`
	Title  string         `gorm:"type:varchar(100)" json:"title, omitempty"`
	Genres pq.StringArray `gorm:"type:varchar(64)" json:"genre, omitempty"`
	Pages  int            `gorm:"type:int" json:"pages, omitempty"`
	Price  float32        `gorm:"type:real" json:"price, omitempty"`
}

//Books comment
type Books []Book

//BookFilter comment
type BookFilter struct {
	Price string `json:"price, omitempty"`
}
