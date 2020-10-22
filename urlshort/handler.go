package urlshort

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var Characters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var PathsToURL = make(map[string]string)

type Url struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)

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
	u.ShortUrl = randSeq(7)

	// place into map
	PathsToURL[u.ShortUrl] = u.LongUrl

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
	longId := PathsToURL[shortId]

	// store retrieved urls into Url struct
	u := Url{
		ShortUrl: shortId,
		LongUrl:  longId,
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
	longId := PathsToURL[shortId]

	if longId == "" {
		http.Error(w, shortId+" not found", http.StatusNotFound)
		return
	}

	// redirect to long_url
	if !strings.HasPrefix(longId, "http://") {
		longId = "http://" + longId
	}
	http.Redirect(w, r, longId, http.StatusMovedPermanently)
}

// generate random string for short_url
func randSeq(n int) string {
	buffer := make([]byte, n)
	for i := range buffer {
		buffer[i] = Characters[rand.Intn(len(Characters))]
	}
	return string(buffer)
}
