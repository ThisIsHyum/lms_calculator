package main

import (
	"log"

	"github.com/ThisIsHyum/lms_calculator/internal/apps/agent"
	"github.com/ThisIsHyum/lms_calculator/internal/config"
)

func main() {
	app := agent.New(config.Config.ComputingPower)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}