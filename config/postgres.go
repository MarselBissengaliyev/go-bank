package config

import (
	"github.com/MarselBissengaliyev/bank/db/postgres"
	"github.com/spf13/viper"
)

func InitPostgresConfig(path string) (*postgres.PostgresConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("postgres")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &postgres.PostgresConfig{
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
		Host:     viper.GetString("host"),
		Port:     viper.GetString("port"),
		DBName:   viper.GetString("dbname"),
		SSLMode:  viper.GetString("sslmode"),
	}, nil
}
