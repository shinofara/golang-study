package controller

import (
	"github.com/shinofara/golang-study/rest/infrastructure"
)

// Base abstract controller
type Base struct {
	DB *infrastructure.DB
}
