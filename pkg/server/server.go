package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"go.uber.org/zap"

	"github.com/hiromaily/go-graphql-server/pkg/server/handler"
	"github.com/hiromaily/go-graphql-server/pkg/server/httpmethod"
)

// Server interface
type Server interface {
	Start() error
	StartTest() (*httptest.Server, error)
	Close()
}

// NewServer returns Server interface
func NewServer(
	logger *zap.Logger,
	schema graphql.Schema,
	method httpmethod.HTTPMethod,
	port int,
	db *sql.DB,
) Server {
	return newServer(
		logger,
		schema,
		method,
		port,
		db,
	)
}

// server object
type server struct {
	logger *zap.Logger
	schema graphql.Schema
	method httpmethod.HTTPMethod
	port   int
	db     *sql.DB
}

// newServer returns server object
func newServer(
	logger *zap.Logger,
	schema graphql.Schema,
	method httpmethod.HTTPMethod,
	port int,
	db *sql.DB,
) *server {
	return &server{
		logger: logger,
		schema: schema,
		method: method,
		port:   port,
		db:     db,
	}
}

// Start starts server
func (s *server) Start() error {
	r := mux.NewRouter()
	if err := handler.GorillaMux(r, s.schema, s.method); err != nil {
		return err
	}

	// server
	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%d", s.port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	s.logger.Info("server is running", zap.Int("port", s.port))
	s.showHelp()

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// for shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done

	s.logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		s.Close()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		s.logger.Error("fatal to call Shutdown():", zap.Error(err))
		return err
	}

	return nil
}

func (s *server) StartTest() (*httptest.Server, error) {
	r := mux.NewRouter()
	if err := handler.GorillaMux(r, s.schema, s.method); err != nil {
		return nil, err
	}
	// test server
	ts := httptest.NewServer(r)

	return ts, nil
}

func (s *server) showHelp() {
	fmt.Printf(`
command:
  curl -g 'http://localhost:8080/graphql?query={user(id:"1"){name}}'
  curl -g 'http://localhost:8080/graphql?query={userList{id,name}}'
  curl -g 'http://localhost:8080/graphql?query=mutation+_{createUser(name:"Tom",age:15,country:"Japan"){id,name,age,country}}'
  curl -g 'http://localhost:8080/graphql?query=mutation+_{updateUser(id:"1",name:"Dummy",age:99,country:"Japan"){id,name,age,country}}'
  curl -g 'http://localhost:8080/graphql?query=mutation+_{deleteUser(id:"2"){id,name,age,country}}'

  curl -g 'http://localhost:8080/graphql?query={company(id:"1"){id,name,country}}'
  curl -g 'http://localhost:8080/graphql?query={companyList{id,name}}'
  curl -g 'http://localhost:8080/graphql?query=mutation+_{createCompany(name:"TechTech",country:"Japan"){id,name,country}}'
  curl -g 'http://localhost:8080/graphql?query=mutation+_{updateCompany(id:"1",name:"TechTechTech"){id,name,country}}'
  curl -g 'http://localhost:8080/graphql?query=mutation+_{deleteCompany(id:"2"){id,name,country}}'

  curl -g 'http://localhost:8080/graphql?query={country(id:"1"){name,name,code}}'
  curl -g 'http://localhost:8080/graphql?query={countryList{id,name}}'

  curl -g 'http://localhost:8080/graphql?query={workHistory(id:"1"){id,company,title}}'
  curl -g 'http://localhost:8080/graphql?query={userWorkHistory(user_id:"1"){id,company,title}}'
  curl -g 'http://localhost:8080/graphql?query={workHistoryList(){id,company,title}}'
  curl -g 'http://localhost:8080/graphql?query=mutation+_{createWorkHistory(user_id:1,company:"Google","backend engineer","tech_ids":[1,2,3],"started_at":"2015/1/1"){id,name,country}}'
  curl -g 'http://localhost:8080/graphql?query=mutation+_{updateWorkHistory(id:1,company:"Google","backend engineer","tech_ids":[1,2,3],"started_at":"2015/1/1"){id,name,country}}'
`)
}

// Close closes dependencies
func (s *server) Close() {
	s.db.Close()
}
