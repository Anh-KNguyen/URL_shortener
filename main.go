package main

import (
	"fmt"
	"time"
	"math/rand"
	"bufio"
	"os"
) 

var characters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
	buffer := make([]byte, n)
	for i := range buffer {
		buffer[i] = characters[rand.Intn(len(characters))] 
	}
	return string(buffer)
}

func main() {
	fmt.Println("URL Shortener")
	url_map := make(map[string] string)

	// read url input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter URL to shorten: ")
	text, _  := reader.ReadString('\n')

	// generate random string
	rand.Seed(time.Now().UnixNano())
	shortenURL := randSeq(10)

	// place into map
	url_map[text] = shortenURL
	fmt.Println("map:", url_map)

}
