package main

import (
	"net/http"

	"github.com/Anh-KNguyen/tinylink/urlshort"
	"github.com/gorilla/mux"
)

// define Rest APIs and Handlers
func defaultMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/links", urlshort.InputHandler).Methods(http.MethodPost)
	r.HandleFunc("/links/{id}", urlshort.OutputHandler).Methods(http.MethodGet)
	r.HandleFunc("/{id}", urlshort.PathHandler)
	r.HandleFunc("/", urlshort.HomeHandler)

	return r
}

func main() {
	// start http server
	mux := defaultMux()
	_ = http.ListenAndServe(":8080", mux)
}
