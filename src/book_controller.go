package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Book declares 'Book' structure
type Book struct {
	Title string
}

// Response declares 'Response' structure
type Response struct {
	NumFound int
	Docs     []Book
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	key := strings.ReplaceAll(keys[0], " ", "+")

	queryURL := "http://openlibrary.org/search.json?q=" + key

	response, err := http.Get(queryURL)

	var booksResponse Response

	if err != nil {
		fmt.Printf("the HTTP request failed with error %s\n", err)

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &booksResponse)
		json.NewEncoder(w).Encode(booksResponse)
	}
}
