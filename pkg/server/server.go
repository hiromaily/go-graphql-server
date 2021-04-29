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
	Clean()
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

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(id:\"1\"){name}}'")
	http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)

	return nil
}

// Clean cleans environment
func (s *server) Clean() {
}

// Close closes dependencies
func (s *server) Close() {
}
