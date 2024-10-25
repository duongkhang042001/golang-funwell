//go:build wireinject
// +build wireinject

package internal

import (
	"core/config"
	"core/internal/entrypoint/http/server"
	"core/pkg/logger"

	"core/internal/infrastructure/persistence/redis"

	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	config.LoadAndParseConfig,
	logger.NewApiLogger,
	redis.NewRedisClient, // Add the RedisClient here
	server.NewApplication,
)

func InitializeApplication(configPath string) (*server.Application, error) {
	wire.Build(providerSet)
	return &server.Application{}, nil
}

func InitializeRedisClient(configPath string) (*redis.RedisClient, error) {
	wire.Build(providerSet)
	return &redis.RedisClient{}, nil
}
