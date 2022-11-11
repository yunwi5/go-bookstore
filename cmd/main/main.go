package main

import (
	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"

	"github.com/yunwi5/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
