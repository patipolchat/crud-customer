package config

import "time"

type (
	Config struct {
		Database DatabaseConfig
		Server   ServerConfig
	}

	DatabaseConfig struct {
		File string `mapstructure:"file" validate:"required"`
	}

	ServerConfig struct {
		Port         int           `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
		LogLevel     string        `mapstructure:"logLevel" validate:"required"`
	}
)
