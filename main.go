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
//
//var storageMap = map[string]func(string) (web.Storage, error){
//	"sqlite":      storage.NewSqliteStorage,
//	"filestorage": storage.NewFileStorage,
//}

func main() {

	flag.Parse()

	var routes *mux.Router

	//storageInit, ok := storageMap[*storageType]
	//if !ok{
	//	log.Fatal("Couldn't set up storage %s, no storage specidied", *storageType)
	//}
	//
	//storage, err := storageInit(*dbPath)
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "Couldn't set up storage"))
	//}
	//routes = web.CreateRoutes(web.NewHandler(storage))

	switch *storageType {
		case "sqlite":
			sqliteStorage, err := storage.NewSqliteStorage(*dbPath)
			if err != nil {
				log.Fatal(errors.Wrap(err, "Couldn't set up storage"))
			}
			routes = web.CreateRoutes(web.NewHandler(sqliteStorage))

		case "filestorage":
			fs, err := storage.NewFileStorage(*dbPath)
			if err != nil {
				log.Fatal(errors.Wrap(err, "Couldn't set up storage"))
			}
			routes = web.CreateRoutes(web.NewHandler(fs))

		default:
			log.Fatal("Couldn't set up a storage, no storage specidied")
			}

	if err := http.ListenAndServe(":8000", routes); err != nil {
		log.Fatal(err)
	}
}
