package sql

import (
	"fmt"
	"time"
)

type SqlPGConfig struct {
	Host                  string
	Port                  uint16 `default:"5432"`
	User                  string
	Password              string
	Name                  string
	SSLMode               string `default:"disable"`
	TimeZone              string `default:"Asia/Bangkok"`
	MaxOpenConns          int    `default:"10"`
	MaxIdleConns          int    `default:"5"`
	ConnMaxLifetimeMinute uint16 `default:"60"`
}

func (c *SqlPGConfig) InitDefaults() {
	if c.Host == "" {
		c.Host = "localhost"
	}
	if c.User == "" {
		c.User = "postgres"
	}
	if c.Name == "" {
		c.Name = "postgres"
	}
	if c.SSLMode == "" {
		c.SSLMode = "disable"
	}
	if c.TimeZone == "" {
		c.TimeZone = "Asia/Bangkok"
	}
	if c.MaxOpenConns <= 0 {
		c.MaxOpenConns = 10
	}
	if c.MaxIdleConns <= 0 {
		c.MaxIdleConns = 5
	}
	if c.ConnMaxLifetimeMinute <= 0 {
		c.ConnMaxLifetimeMinute = 60
	}
}

func (c *SqlPGConfig) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.Host,
		c.User,
		c.Password,
		c.Name,
		c.Port,
		c.SSLMode,
		c.TimeZone,
	)
}
func (c *SqlPGConfig) GetConnMaxLifetime() time.Duration {
	return time.Duration(c.ConnMaxLifetimeMinute) * time.Minute
}

type SqlMySQLConfig struct {
	Host                  string
	Port                  uint16 `default:"3306"`
	User                  string
	Password              string
	Name                  string
	Charset               string `default:"utf8mb4"`
	Collation             string `default:"utf8mb4_general_ci"`
	MaxOpenConns          int    `default:"10"`
	MaxIdleConns          int    `default:"5"`
	ConnMaxLifetimeMinute uint16 `default:"60"`
}

func (c *SqlMySQLConfig) InitDefaults() {
	if c.Host == "" {
		c.Host = "localhost"
	}
	if c.User == "" {
		c.User = "root"
	}
	if c.Name == "" {
		c.Name = "mysql"
	}
	if c.Charset == "" {
		c.Charset = "utf8mb4"
	}
	if c.Collation == "" {
		c.Collation = "utf8mb4_general_ci"
	}
	if c.MaxOpenConns <= 0 {
		c.MaxOpenConns = 10
	}
	if c.MaxIdleConns <= 0 {
		c.MaxIdleConns = 5
	}
	if c.ConnMaxLifetimeMinute <= 0 {
		c.ConnMaxLifetimeMinute = 60
	}
}
func (c *SqlMySQLConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&parseTime=true&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		c.Charset,
		c.Collation,
	)
}
func (c *SqlMySQLConfig) GetConnMaxLifetime() time.Duration {
	return time.Duration(c.ConnMaxLifetimeMinute) * time.Minute
}
