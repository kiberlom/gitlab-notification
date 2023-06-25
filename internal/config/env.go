package config

import "github.com/spf13/viper"

type ConfigENV struct {
	TelegramBotToken   string
	TelegramGroupAdmin string
	ServerHTTPPort     int
}

func NewENV() *ConfigENV {
	v := viper.New()
	// _ = v.ReadInConfig()
	v.AutomaticEnv()

	return &ConfigENV{
		TelegramBotToken:   v.GetString("TELEGRAM_BOT_TOKEN"),
		TelegramGroupAdmin: v.GetString("TELEGRAM_GROUP_ADMIN"),
		ServerHTTPPort:     v.GetInt("SERVER_HTTP_PORT"),
	}
}
