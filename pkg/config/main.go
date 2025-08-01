package cf

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig[T any]() *T {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/app")
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
