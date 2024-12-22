package main

import (
	"fmt"

	application "github.com/ThisIsHyum/lms_calculator/internal/app"
)

func main() {
	app := application.New()
	err := app.Run()
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
}
