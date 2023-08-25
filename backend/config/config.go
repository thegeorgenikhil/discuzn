package config

import (
	"github.com/spf13/viper"
)

type ConfigName string

const (
	PORT ConfigName = "PORT"
)

type AppConfigI interface {
	LoadConfig()
	GetPort() string
}

type AppConfig struct {
	port string
}

func NewAppConfig() AppConfigI {
	return &AppConfig{}
}

func (a *AppConfig) LoadConfig() {
	viper.AutomaticEnv()
	viper.SetConfigName(".env")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../")
	viper.SetConfigType("env")
	viper.ReadInConfig()

	a.port = mustGetString(PORT)
}

func (a *AppConfig) GetPort() string {
	return a.port
}
