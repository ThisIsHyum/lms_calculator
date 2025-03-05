package main

import (
	"log"

	"github.com/ThisIsHyum/lms_calculator/internal/apps/orchestrator"
)

func main() {
	app := orchestrator.New()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}