package config

import (
	"github.com/goravel/framework/contracts/application"
	"github.com/goravel/framework/facades"
)

func App() map[string]any {
	return map[string]any{
		"name":     facades.Config().Env("APP_NAME", "Goravel"),
		"env":      facades.Config().Env("APP_ENV", "local"),
		"debug":    facades.Config().Env("APP_DEBUG", false),
		"timezone": facades.Config().Env("APP_TIMEZONE", "UTC"),
		"url":      facades.Config().Env("APP_URL", "http://localhost"),
		"port":     facades.Config().Env("APP_PORT", "3000"),
		"key":      facades.Config().Env("APP_KEY", ""),
		"providers": []application.ServiceProvider{
			// Add your service providers here
		},
	}
}
