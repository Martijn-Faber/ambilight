package config

import (
	"github.com/spf13/viper"
)

type Config struct{
	NumLedsX int `mapstructure:"NUM_LEDS_X"`
	NumLedsY int `mapstructure:"NUM_LEDS_Y"`
	RefTimeMs int `mapstructure:"REFRESH_TIME_IN_MS"`
	WledIp string `mapstruture:"WLED_IP"`
	WledPort int `mapstructure:"WLED_PORT"`
}

func Load() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	// TODO: print error to set the .env vars 
	if err != nil {
		panic(err)
	}

	return
}