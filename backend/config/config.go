package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigName string

const (
	PORT       ConfigName = "PORT"
	DBNAME     ConfigName = "DBNAME"
	DBUSERNAME ConfigName = "DBUSERNAME"
	DBPASS     ConfigName = "DBPASS"
	DBHOST     ConfigName = "DBHOST"
	DBPORT     ConfigName = "DBPORT"
)

type AppConfigI interface {
	LoadConfig()
	GetPort() string
	GetDbURI() string
}

type AppConfig struct {
	port       string
	dbName     string
	dbUsername string
	dbPassword string
	dbHost     string
	dbPort     string
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

	a.dbName = mustGetString(DBNAME)
	a.dbUsername = mustGetString(DBUSERNAME)
	a.dbPassword = mustGetString(DBPASS)
	a.dbHost = mustGetString(DBHOST)
	a.dbPort = mustGetString(DBPORT)

}

func (a *AppConfig) GetPort() string {
	return a.port
}

func (a *AppConfig) GetDbURI() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", a.dbUsername, a.dbPassword, a.dbHost, a.dbPort, a.dbName)
}
