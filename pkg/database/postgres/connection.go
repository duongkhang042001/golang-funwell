package postgres

import (
	"core/config"
	"core/pkg/logger"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	gormLogger "gorm.io/gorm/logger"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120 * time.Minute
	maxIdleConns    = 30
	connMaxIdleTime = 20 * time.Minute
)

func NewPostgresDB(cfg *config.Config, logger logger.Logger) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		cfg.Postgres.PostgresqlHost, cfg.Postgres.PostgresqlUser, cfg.Postgres.PostgresqlPassword, cfg.Postgres.PostgresqlDbname, cfg.Postgres.PostgresqlPort)

	gormConfig := &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		logger.Errorf("Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorf("Failed to get SQL DB from Gorm: %v", err)
		return nil, err
	}
	defer sqlDB.Close()

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxIdleTime(connMaxIdleTime)

	logger.Info("Successfully connected to PostgreSQL database")

	return db, nil
}
