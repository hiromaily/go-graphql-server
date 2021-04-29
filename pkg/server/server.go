package server

import "go.uber.org/zap"

// Server interface
type Server interface {
	Start() error
	Clean()
	Close()
}

// NewServer returns Server interface
func NewServer(
	logger *zap.Logger,
	port int,
) Server {
	return newServer(
		logger,
		port,
	)
}

// server object
type server struct {
	logger *zap.Logger
	port   int
}

// NewBook is to return book object
func newServer(
	logger *zap.Logger,
	port int,
) *server {
	return &server{
		logger: logger,
		port:   port,
	}
}

// Start starts server
func (s *server) Start() error {
	return nil
}

// Clean cleans environment
func (s *server) Clean() {
}

// Clean closes dependencies
func (s *server) Close() {
}
