package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
