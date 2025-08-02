package db

import "time"

type DBConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Dbname          string
	SSLMode         string
	TimeZone        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}
