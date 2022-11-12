package main

import (
	"github.com/firehead666/infotecs-go-test-task/client/internal/app"
	"github.com/firehead666/infotecs-go-test-task/client/internal/config"
	"log"
)

func main() {
	cfg := config.GetConfig()

	a, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}

	a.Run()
}
