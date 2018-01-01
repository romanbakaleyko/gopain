package storage

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

//BookFilter comment
type BookFilter struct {
	Price string `json:"price, omitempty"`
}
