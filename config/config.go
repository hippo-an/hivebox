package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var Config Cfg

type Cfg struct {
	Server struct {
		Port    string `mapstructure:"port"`
		Version string `mapstructure:"version"`
	} `mapstructure:"server"`
	ForecastServiceKey string `mapstructure:"forecast_service_key"`
}

func InitConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("fatal error reading config file: %w", err))
	}

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	if err := viper.MergeInConfig(); err != nil {
		log.Println("No .env file found, relying on environment variables")

		viper.BindEnv("forecast_service_key", "FORECAST_SERVICE_KEY")
	}

	viper.SetEnvPrefix("HIVEBOX")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var appConfig Cfg

	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Fatal(fmt.Errorf("fatal error to unmarshal config struct: %w", err))
	}

	Config = appConfig
}
