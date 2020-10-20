package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var characters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var pathsToURL = make(map[string]string)

type Url struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

// define Rest APIs and Handlers
func defaultMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/links", InputHandler).Methods(http.MethodPost)
	r.HandleFunc("/links/{id}", OutputHandler).Methods(http.MethodGet)
	r.HandleFunc("/{id}", PathHandler).Methods(http.MethodGet)
	return r
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func InputHandler(w http.ResponseWriter, r *http.Request) {
	// define new Url struct
	var u Url

	// decode the request body into Url struct
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// generate random short_url
	rand.Seed(time.Now().UnixNano())
	u.ShortUrl = randSeq(10)

	// place into map
	pathsToURL[u.ShortUrl] = u.LongUrl

	// encode struct into json string representation
	err = json.NewEncoder(w).Encode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func OutputHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve route
	vars := mux.Vars(r)
	shortId := vars["id"]
	longId := pathsToURL[shortId]

	// store retrieved urls into Url struct
	u := Url{
		ShortUrl: shortId, 
		LongUrl: longId,
	}

	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PathHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve route
	vars := mux.Vars(r)
	shortId := vars["id"]
	longId := pathsToURL[shortId]

	http.Redirect(w, r, longId, http.StatusMovedPermanently)
	return
}

// generate random string for short_url
func randSeq(n int) string {
	buffer := make([]byte, n)
	for i := range buffer {
		buffer[i] = characters[rand.Intn(len(characters))]
	}
	return string(buffer)
}

func main() {
	// start http server
	mux := defaultMux()
	_ = http.ListenAndServe(":8080", mux)
	
}
