package application

import (
	"net/http"

	"github.com/ThisIsHyum/lms_calculator/internal/handlers"
)

type App struct {
	Port string
}

func New() App {
	return App{
		Port: ":80",
	}
}

func (a App) Run() {
	http.HandleFunc("/api/v1/calculate", handlers.Calculate)
	http.ListenAndServe(a.Port, nil)
}
