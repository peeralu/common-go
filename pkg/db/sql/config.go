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
