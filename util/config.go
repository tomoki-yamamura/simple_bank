package util

import "github.com/spf13/viper"

type Config struct {
	DB_DRIVER string `mapstructure:"db_driver"`
	DB_SOURCE string `mapstructure:"db_source"`
	SERVER_ADDRESS string `mapstructure:"server_address"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}