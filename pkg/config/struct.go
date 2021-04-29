package config

import "github.com/hiromaily/go-graphql-server/pkg/server/httpmethod"

// Root is root config
type Root struct {
	Server *Server `toml:"server" validate:"required"`
	Logger *Logger `toml:"logger" validate:"required"`
}

// Server is server information
type Server struct {
	Port       int                   `toml:"port" validate:"required"`
	HTTPMethod httpmethod.HTTPMethod `toml:"http_method" validate:"oneof=GET POST"`
}

// Logger is zap logger property
type Logger struct {
	Service      string `toml:"service" validate:"required"`
	Env          string `toml:"env" validate:"oneof=dev prod custom"`
	Level        string `toml:"level" validate:"required"`
	IsStackTrace bool   `toml:"is_stacktrace"`
}
