package redis

import (
	"context"
	"core/config"
	"core/pkg/logger"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewRedisClient(cfg *config.Config, logger logger.Logger) *redis.Client {
	redisHost := cfg.Redis.RedisHost

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: cfg.Redis.MinIdleConns,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Info("Failed to connect to Redis")
	}
	logger.Info("Connected to Redis")

	return client
}
