package util

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig[T any](fileName string, fileType string, paths []string) *T {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var cf T
	err = viper.Unmarshal(&cf)
	if err != nil {
		log.Panicf("[InitConfig] Failed to unmarshal config: %v", err)
	}

	return &cf
}
