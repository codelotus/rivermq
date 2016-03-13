package main

import (
	"log"
	"net/http"

	"github.com/codelotus/rivermq/route"
)

func main() {
	router := route.NewRouter()
	log.Println("Started, listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
