package controller

import (
	"net/http"
	"fmt"
)

// Index controller.
type Index struct {
	Base
}

func (t *Index) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
