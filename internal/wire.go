//go:build wireinject
// +build wireinject

package internal

import (
	"core/config"
	"core/internal/entrypoint/http/server"
	"core/pkg/logger"


	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	config.LoadAndParseConfig,
	logger.NewApiLogger,
	server.NewApplication,
)

func InitializeApplication(configPath string) (*server.Application, error) {
	wire.Build(providerSet)
	return &server.Application{}, nil
}