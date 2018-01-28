package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/romanbakaleyko/gopain/api/web"
	"github.com/romanbakaleyko/gopain/storage"
	log "github.com/sirupsen/logrus"
)

var (
	storageType = flag.String("storage", "", "--storage slite")
	dbPath      = flag.String("path", "", "storage/storage.db")
)

func main() {

	flag.Parse()

	var routes *mux.Router

	switch *storageType {
		case "sqlite":
			sqliteStorage, err := storage.NewSqliteStorage(*dbPath)
			if err != nil {
				log.Fatal(errors.Wrap(err, "couldn't set up sqlite storage"))
			}
			routes = web.CreateRoutes(web.NewHandler(sqliteStorage))

		case "filestorage":
			fs, err := storage.NewFileStorage(*dbPath)
			if err != nil {
				log.Fatal(errors.Wrap(err, "couldn't set up file storage"))
			}
			routes = web.CreateRoutes(web.NewHandler(fs))

		default:
			log.Fatal("couldn't set up a storage, no storage specified")
			}

	if err := http.ListenAndServe(":8000", routes); err != nil {
		log.Fatal(err)
	}
}
