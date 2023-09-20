package config

import (
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type Config struct {
	Api                 ApiConfig
	Db                  DatabaseConfig
	HttpClient          HttpClientConfig
	PortalTransparencia PortalTransparenciaConfig
}

type ApiConfig struct {
	Port string
}

type DatabaseConfig struct {
	Port     string
	Host     string
	Username string
	Password string
	Database string
}

type HttpClientConfig struct {
	BaseUrl string
}

type PortalTransparenciaConfig struct {
	Key string
}

var cfg *Config

func init() {
	viper.SetDefault("Api.Port", "9000")
	viper.SetDefault("Database.Host", "localhost")
	viper.SetDefault("Database.Port", "5432")
	viper.SetDefault("HttpClient.BaseUrl", "https://api.portaldatransparencia.gov.br/api-de-dados/")
	viper.SetDefault("PortalTransparencia.Key", "0c0c0ca0a6f1d72d")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(Config)

	cfg.Api = ApiConfig{
		Port: viper.GetString("Api.Port"),
	}

	cfg.Db = DatabaseConfig{
		Host:     viper.GetString("Database.Host"),
		Port:     viper.GetString("Database.Port"),
		Username: viper.GetString("Database.Username"),
		Password: viper.GetString("Database.Password"),
		Database: viper.GetString("Database.Name"),
	}

	cfg.HttpClient = HttpClientConfig{
		BaseUrl: viper.GetString("HttpClient.BaseUrl"),
	}

	cfg.PortalTransparencia = PortalTransparenciaConfig{
		Key: viper.GetString("PortalTransparencia.Key"),
	}

	return nil
}

func GetDb() DatabaseConfig {
	return cfg.Db
}

func GetDbConnection() mysql.Config {
	c := mysql.Config{
		User:   cfg.Db.Username,
		Passwd: cfg.Db.Password,
		Net:    "tcp",
		Addr:   cfg.Db.Host,
		DBName: cfg.Db.Database,
	}

	return c
}

func GetServerPort() string {
	return cfg.Api.Port
}

func GetBaseUrl() string {
	return cfg.HttpClient.BaseUrl
}

func GetChaveApi() string {
	return cfg.PortalTransparencia.Key
}
