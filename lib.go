package main

import (
	secretManager "github.com/peeralu/common-go/pkg/aws/secret_manager"
	"github.com/peeralu/common-go/pkg/config"
	"github.com/peeralu/common-go/pkg/db/sql"
	"github.com/peeralu/common-go/pkg/time"
)

var (
	GetSecretFromAWSSecretManager = secretManager.GetSecretFromAWSSecretManager

	NewTimeZone     = time.NewTimeZone
	NewPostgreSQLDB = sql.NewPostgreSQLDB
)

func InitConfig[T any]() *T {
	return config.InitConfig[T]()
}
