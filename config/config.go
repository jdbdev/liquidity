package config

import "time"

type AppConfig struct {
	TickerSettings TickerSettings
	DBSettings     DBSettings
}

type DBSettings struct {
}

type TickerSettings struct {
	Intervale time.Duration
	APIKey    string
	BaseURL   string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}
