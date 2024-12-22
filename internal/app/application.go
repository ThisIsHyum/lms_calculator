package application

import (
	"net/http"

	"github.com/ThisIsHyum/lms_calculator/internal/config"
	"github.com/ThisIsHyum/lms_calculator/internal/handlers"
)

type App struct {
	Ip   string
	Port string
}

func New() App {
	return App{
		Ip:   *config.Ip,
		Port: *config.Port,
	}
}

func (a App) Run() error {
	http.HandleFunc("/api/v1/calculate", handlers.Calculate)
	return http.ListenAndServe(a.Ip+":"+a.Port, nil)
}
