package main

import (
	"github.com/go-ewallet/client/internal/app"
	"github.com/go-ewallet/client/internal/config"
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
