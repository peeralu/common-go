package sql

import (
	"fmt"

	"github.com/peeralu/common-go/pkg/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgreSQLDB(config db.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Host,
		config.User,
		config.Password,
		config.Dbname,
		config.Port,
		config.SSLMode,
		config.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)

	return db
}
