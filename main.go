package main

import (
	"log"
	"net/http"

	"github.com/romanbakaleyko/gopain/api/web"
)

func main() {
	routes := web.CreateRoutes()

	if err := http.ListenAndServe(":8000", routes); err != nil {
		log.Fatal(err)
	}
}
