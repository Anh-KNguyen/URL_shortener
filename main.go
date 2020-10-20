package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/Anh-KNguyen/URL_shortener/urlshort"
	"github.com/gorilla/mux"
)

var characters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Url struct {
	LongUrl string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

// define Rest APIs and Handlers
func defaultMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/links", InputHandler).Methods(http.MethodPost)
	r.HandleFunc("/links/{id}", OutputHandler).Methods(http.MethodGet)
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


}

func OutputHandler(w http.ResponseWriter, r *http.Request) {

}

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
	go http.ListenAndServe(":8080", mux)

	// Build MapHandler using mux as fallback
	pathsToURL := make(map[string]string)
	mapHandler := urlshort.MapHandler(pathsToURL, mux)

	// read url input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter URL to shorten: ")
	text, _ := reader.ReadString('\n')

	// generate random string
	rand.Seed(time.Now().UnixNano())
	shortenURL := randSeq(10)

	// place into map
	pathsToURL[text] = shortenURL
	fmt.Println("map:", pathsToURL)
	fmt.Println(mapHandler)

}
