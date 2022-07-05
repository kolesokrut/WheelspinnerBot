package config

import "github.com/spf13/viper"

type Config struct {
	TelegramToken string `mapstructure:"TELEGRAM-TOKEN"`
}

func LoadConfig(name string) (config Config, err error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
