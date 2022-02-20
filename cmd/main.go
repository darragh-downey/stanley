package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/darragh-downey/stanley/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)

	http.Handle("/", r)
}
