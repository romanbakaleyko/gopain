package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type Book struct {
	id    string
	title string
	genre []string
	pages int
	price float32
}

func main() {

	book := Book{id: uuid.NewV4().String(), title: "Test"}
	fmt.Println(book)

}
