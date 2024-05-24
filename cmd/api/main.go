package main

import (
	"core/config"
	"core/internal/infrastructure/server"
	"core/pkg/logger"
	"log"
)

func main() {
	cfg, err := config.LoadAndParseConfig("./config/config-local")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()

	app := server.NewApplication(cfg, logger)

	logger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	if err := app.Start(); err != nil {
		logger.Error(err)
	}
}

// https://sarathsp06.medium.com/domain-driven-design-with-go-be3066ae213c
// https://dev.to/ayoubzulfiqar/go-the-ultimate-folder-structure-6gj
