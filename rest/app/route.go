package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shinofara/golang-study/rest/app/middleware"
	"github.com/shinofara/golang-study/rest/infrastructure"
	"github.com/shinofara/golang-study/rest/app/controller"
)

// NewMux create the handler.
func NewMux() http.Handler {
	r := mux.NewRouter()

	db, err := infrastructure.NewDB()
	if err != nil {
		log.Panic(err)
	}

	r.Use(middleware.Authenticate(db))
	r.Use(middleware.RequestLogger)

	base := controller.Base{
		DB: db,
	}
	transaction := controller.TransactionController{Base: base}
	index := controller.Index{Base: base}
	r.HandleFunc("/", index.Index).Methods(http.MethodGet)
	r.HandleFunc("/transactions", transaction.List).Methods(http.MethodGet)
	r.HandleFunc("/transactions/{id:[0-9]+}", transaction.Show).Methods(http.MethodGet)
	r.HandleFunc("/transactions", transaction.Create).Methods(http.MethodPost)
	r.HandleFunc("/transactions/{id:[0-9]+}", transaction.Delete).Methods(http.MethodDelete)

	return r
}
