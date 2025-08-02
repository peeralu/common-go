package sql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgreSQLDB(config SqlPGConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Warn),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get database instance: %v", err))
	}

	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(config.GetConnMaxLifetime())

	return db
}
