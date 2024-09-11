package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

var C Config

const (
	configName = "config"
	configType = "yaml"
	configPath = "."
)

func InitConfig() {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		message := fmt.Errorf("failed to read config file:%v", err)
		panic(message)
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		message := fmt.Errorf("failed to unmarshal config:%v", err)
		panic(message)
	}
}
