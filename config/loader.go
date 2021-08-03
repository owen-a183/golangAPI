package config

import (
	"github.com/spf13/viper"
)

func Load() (*Config, error) {
	viper.SetConfigName(".config") // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")       // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
