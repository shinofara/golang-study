package main

import (
	"net/http"
	"fmt"
	"github.com/shinofara/golang-study/fizzbuzz"
	"strconv"
)

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/nocontent", noContentHandler)
  http.HandleFunc("/json", jsonHandler)
  http.HandleFunc("/fizzbuzz", fizzBuzzHandler)
  http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
    fmt.Fprint(w, "1")
}

func noContentHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(204)
}

func jsonHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, `{"year": "2019", "status": 200}`)
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	if v == nil {
		return
	}

	n := v.Get("n")
	ni, err := strconv.Atoi(n)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprintf(w, `{"value":"%s"}`, fizzbuzz.Run(uint32(ni)))
	return
}
