package main

import (
	"exercise/src"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", src.Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
