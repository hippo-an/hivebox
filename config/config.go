package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var AppConfig HiveBoxConfig
var AppSecret HiveBoxSecret

type HiveBoxConfig struct {
	Application struct {
		Port    string `yaml:"port"`
		Version string `yaml:"version"`
	}
}

type HiveBoxSecret struct {
	ForecastServiceKey string `yaml:"forecastServiceKey"`
}

func InitConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("fatal error reading config file: %w", err))
	}

	var appConfig HiveBoxConfig

	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Fatal(fmt.Errorf("fatal error to unmarshal config struct: %w", err))
	}

	AppConfig = appConfig
}

func InitSecret() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")

	viper.SetConfigName("secret")
	viper.SetConfigType("yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("fatal error reading secret file: %w", err))
	}

	var appSecret HiveBoxSecret

	if err := viper.Unmarshal(&appSecret); err != nil {
		log.Fatal(fmt.Errorf("fatal error to unmarshal secret struct: %w", err))
	}

	AppSecret = appSecret
}
