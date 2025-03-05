package orchestrator

import (
	"github.com/ThisIsHyum/lms_calculator/internal/apps/orchestrator/handlers"
	"github.com/ThisIsHyum/lms_calculator/internal/config"
	"github.com/labstack/echo/v4"
)

type App struct {
	Ip   string
	Port string
}

func New() App {
	return App{
		Port: config.Config.Port,
	}
}

func (a App) Run() error {
	e := echo.New()

	e.POST("/api/v1/calculate", handlers.AddExpression)
	e.GET("/api/v1/expressions", handlers.GetExpressions)
	e.GET("/api/v1/expressions/:id", handlers.GetExpression)

	e.GET("/internal/task", handlers.GetTask)
	e.POST("/internal/task", handlers.SendResult)


	return e.Start(":" + a.Port)
}
