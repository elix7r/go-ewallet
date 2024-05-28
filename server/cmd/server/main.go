package main

import (
	"github.com/go-ewallet/server/internal/app"
	"github.com/go-ewallet/server/internal/config"
	"github.com/go-ewallet/server/pkg/logging"
	"log"
)

func main() {
	log.Println("initializing configuration...")
	cfg := config.GetConfig()

	log.Println("initializing logger...")
	logger := logging.GetLogger(cfg.AppConfig.LogLevel)

	a, err := app.NewApp(cfg, &logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("starting to run server application...")
	a.Run()
}
