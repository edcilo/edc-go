package edc

import (
	"github.com/gofiber/fiber/v2/log"
)

var Config = Configuration{}

func ConfigSetup(envfiles ...string) *Configuration {
	log.Info("Setting up configuration")

	env := Env{}
	env.Load(envfiles...)

	Config = Configuration{
		App: ConfigurationApp{
			Name:        env.Get("APP_NAME", "app"),
			Description: env.Get("APP_DESCRIPTION", ""),
			Version:     env.Get("APP_VERSION", "0.0.1"),
			Host:        env.Get("APP_HOST", "localhost"),
			Port:        env.GetInt("APP_PORT", 8080),
		},
		DB: ConfigurationDB{
			Engine:   DBEngine(env.Get("DB_ENGINE", string(Postgres))),
			Host:     env.Get("DB_HOST", "localhost"),
			Port:     env.GetInt("DB_PORT", 5432),
			Database: env.Get("DB_DATABASE", "postgres"),
			User:     env.Get("DB_USER", "postgres"),
			Password: env.Get("DB_PASSWORD", "postgres"),
		},
		Cache: ConfigurationCache{
			Engine:   CacheEngine(env.Get("CACHE_ENGINE", string(Redis))),
			Host:     env.Get("CACHE_HOST", "localhost"),
			Port:     env.GetInt("CACHE_PORT", 6379),
			Database: env.GetInt("CACHE_DATABASE", 0),
			User:     env.Get("CACHE_USER", "default"),
			Password: env.Get("CACHE_PASSWORD", ""),
		},
	}

	return &Config
}
