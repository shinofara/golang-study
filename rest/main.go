package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/shinofara/golang-study/rest/app"
)

func main() {
	if err := http.ListenAndServe(":8888", app.NewMux()); err != nil {
		log.Println(err)
	}
}
