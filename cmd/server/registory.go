package main

import (
	"go.uber.org/zap"

	"github.com/hiromaily/go-graphql-server/pkg/config"
	"github.com/hiromaily/go-graphql-server/pkg/logger"
	"github.com/hiromaily/go-graphql-server/pkg/server"
)

// Registry interface
type Registry interface {
	NewServer() server.Server
}

type registry struct {
	conf   *config.Root
	logger *zap.Logger
}

// NewRegistry is to register regstry interface
func NewRegistry(conf *config.Root) Registry {
	return &registry{conf: conf}
}

// NewServer registers for Server interface
func (r *registry) NewServer() server.Server {
	return server.NewServer(
		r.newLogger(),
		r.conf.Server.Port,
	)
}

func (r *registry) newLogger() *zap.Logger {
	if r.logger == nil {
		r.logger = logger.NewZapLogger(r.conf.Logger)
	}
	return r.logger
}
