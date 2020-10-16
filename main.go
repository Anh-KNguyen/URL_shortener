package main

import (
	"fmt"
	"time"
	"math/rand"
) 

var characters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	buffer := make([]byte, n)
	for i := range buffer {
		buffer[i] = characters[rand.Intn(len(characters))] 
	}
	return string(buffer)
}

func shortenURL(url string) {

}

func main() {
	fmt.Println("URL Shortener")
	url_map := make(map[string] string)

	rand.Seed(time.Now().UnixNano())
	fmt.Println(randSeq(10))

	fmt.Println("map:", url_map)
}
