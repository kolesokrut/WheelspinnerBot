package config

import "github.com/spf13/viper"

type Api struct {
	OpenWeather string `mapstructure:"openweather"`
	Soundcloud  string `mapstructure:"soundcloud"`
	Tiktok      string `mapstructure:"tiktok"`
	Youtube     string `mapstructure:"youtube"`
}

type DateBase struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Protocol string `mapstructure:"protocol"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Driver   string `mapstructure:"driver"`
}

type Config struct {
	TelegramToken string   `mapstructure:"TELEGRAM-TOKEN"`
	DB            DateBase `mapstructure:"db"`
	Api           Api      `mapstructure:"api"`
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
