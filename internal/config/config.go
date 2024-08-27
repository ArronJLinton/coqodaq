package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// InitConfig fetches the configuration from the .env file
func InitConfig() Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetDefault("db_url", "")
	viper.SetDefault("port", "8080")

	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			log.Fatal("Error reading config file - ", err)
		} else {
			// Config file was found but another error was produced
			log.Fatal("Error reading config file - ", err)
		}
	}

	var c Config
	viper.Unmarshal(&c)

	return c
}
