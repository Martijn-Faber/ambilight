package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	NumLedsX      int    `mapstructure:"NUM_LEDS_X"`
	NumLedsY      int    `mapstructure:"NUM_LEDS_Y"`
	RefTimeMs     int    `mapstructure:"REFRESH_TIME_IN_MS"`
	WledIp        string `mapstructure:"WLED_IP"`
	WledPort      int    `mapstructure:"WLED_PORT"`
	InterpolSteps int    `mapstructure:"INTERPOLATE_STEPS"`
	RCorrection float32 `mapstructure:"R_CORRECTION"`
	GCorrection float32 `mapstructure:"G_CORRECTION"`
	BCorrection float32 `mapstructure:"B_CORRECTION"`
}

func Load() (config Config, err error) {
	wd, err := os.Getwd()

	if err != nil {
		return
	}

	viper.AddConfigPath(wd)
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
