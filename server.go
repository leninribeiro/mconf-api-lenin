package main

import (
	"log"
	"net/http"
	"os"

	"github.com/leninribeiro/mconf-api-lenin/src"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", src.GetBooks)

	n := negroni.Classic()
	n.UseHandler(mux)

	log.Println("Server started on port", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), n)
}
