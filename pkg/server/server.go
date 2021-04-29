package server

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"go.uber.org/zap"

	"github.com/hiromaily/go-graphql-server/pkg/server/handler"
)

// Server interface
type Server interface {
	Start() error
	Close()
}

// NewServer returns Server interface
func NewServer(
	logger *zap.Logger,
	schema graphql.Schema,
	port int,
) Server {
	return newServer(
		logger,
		schema,
		port,
	)
}

// server object
type server struct {
	logger *zap.Logger
	schema graphql.Schema
	port   int
}

// newServer returns server object
func newServer(
	logger *zap.Logger,
	schema graphql.Schema,
	port int,
) *server {
	return &server{
		logger: logger,
		schema: schema,
		port:   port,
	}
}

// Start starts server
func (s *server) Start() error {
	handler.Initialize(s.schema)

	s.logger.Info("server is running", zap.Int("port", s.port))
	fmt.Printf(`
command:
  curl -g 'http://localhost:%d/graphql?query={user(id:"1"){name}}'
  curl -g 'http://localhost:%d/graphql?query={userList{id,name}}'
`, s.port, s.port)
	http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)

	return nil
}

// Close closes dependencies
func (s *server) Close() {
}
