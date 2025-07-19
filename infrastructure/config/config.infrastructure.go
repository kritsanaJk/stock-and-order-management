package config

import (
	"github.com/spf13/viper"
)

var AppConfig *ServiceConfig

type ServiceConfig struct {
	Mode             string `mapstructure:"MODE"`
	Port             string `mapstructure:"PORT"`
	DatabaseLocation string `mapstructure:"DATABASE_LOCATION"`
}

func LoadConfig(path string) (config *ServiceConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err = viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return
}
