package main

import (
	"net/http"
	"fmt"
)

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/nocontent", noContentHandler)
  http.HandleFunc("/json", jsonHandler)
  http.HandleFunc("/fizzbuzz", fizzBuzzHandler)
  http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "1")
}

func noContentHandler(_ http.ResponseWriter, _ *http.Request) {
}

func jsonHandler(_ http.ResponseWriter, _ *http.Request) {
}

func fizzBuzzHandler(_ http.ResponseWriter, _ *http.Request) {
}