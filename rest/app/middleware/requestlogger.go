package middleware

import (
	"log"
	"net/http"
)

func RequestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%+v\n", r)
		handler.ServeHTTP(w, r)
	})
}
