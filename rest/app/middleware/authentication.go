package middleware

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/shinofara/golang-study/rest/infrastructure"
	"log"
)

type UserIDKeyType string

const UserIDKey UserIDKeyType = "USER_ID"

// Authenticate verifies API key and set User ID to the request context.
func Authenticate(db *infrastructure.DB) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			apiKey := r.Header.Get("apikey")



			var userID int
			if err := db.Open().QueryRowContext(
				ctx,
				"select id from users where api_key=?",
				apiKey,
			).Scan(&userID); err != nil {
				if err == sql.ErrNoRows {
					http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			log.Printf("apikey is %s", apiKey)
			handler.ServeHTTP(w, r.WithContext(context.WithValue(ctx, UserIDKey, userID)))
		})
	}
}
