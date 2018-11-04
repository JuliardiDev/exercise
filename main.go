package main

import (
	"log"
	"net/http"

	"github.com/exercise/src"
)

func main() {
	http.HandleFunc("/", src.Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
