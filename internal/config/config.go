package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var AppConfig Config

type Config struct {
	ServerAddr    string `mapstructure:"SERVER_ADDR"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig() {
	// Search config in home directory with name ".filepath" (without extension).
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
