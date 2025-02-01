package main

import (
	"context"
	"core/config"
	"core/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.LoadAndParseConfig("./config/config-local")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	logger := logger.NewApiLogger(cfg)

	logger.InitLogger()

	logger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)

	logger.Infof("Worker is running")

	_, cancel := context.WithCancel(context.Background())

	terminateSignals := make(chan os.Signal, 1)

	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case s := <-terminateSignals:
			logger.Info("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME :", s)
			cancel()
			stop = true
		default:
			time.Sleep(1 * time.Second)
		}
		
	}

	time.Sleep(5 * time.Second) // wait for all consumers to finish processing
}
