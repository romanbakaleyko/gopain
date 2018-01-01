package main

import (
	"net/http"

	"github.com/romanbakaleyko/gopain/api/web"
	log "github.com/sirupsen/logrus"
)

func main() {

	routes := web.CreateRoutes()

	if err := http.ListenAndServe(":8000", routes); err != nil {
		log.Fatal(err)
	}
}
