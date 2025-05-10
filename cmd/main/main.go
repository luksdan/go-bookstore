package main

import (
	"log"
	"net/http"

	"github.com/luksdan/go-bookstore/pkg/routes"
)

func main() {
	routes.RegisterBookStoreRoutes()
	log.Fatal(http.ListenAndServe("localhost:9010", nil))
}
