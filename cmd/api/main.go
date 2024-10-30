package main

import (
	"core/config"
	"core/internal/interfaces/http/server"
	"core/pkg/database/postgres"
	"core/pkg/database/redis"
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

	redisClient := redis.NewRedisClient(cfg, logger)
	defer redisClient.Close()

	postgresClient, err := postgres.NewPostgresDB(cfg, logger)
	if err != nil {
		logger.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	app := server.NewApplication(cfg, logger, postgresClient, redisClient)

	logger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	if err := app.Start(); err != nil {
		logger.Error(err)
	}
}

// https://sarathsp06.medium.com/domain-driven-design-with-go-be3066ae213c
// https://dev.to/ayoubzulfiqar/go-the-ultimate-folder-structure-6gj
