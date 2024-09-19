package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

const (
	configName = "config"
	configType = "yaml"
	configPath = "."
)

var (
	c    Config
	once sync.Once
)

type Config struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
		viper.AddConfigPath(configPath)

		err := viper.ReadInConfig()
		if err != nil {
			message := fmt.Errorf("failed to read config file:%v", err)
			log.Fatalln(message)
		}

		err = viper.Unmarshal(&c)
		if err != nil {
			message := fmt.Errorf("failed to unmarshal config:%v", err)
			log.Fatalln(message)
		}
	})

	return &c
}
