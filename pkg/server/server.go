package server

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"go.uber.org/zap"
	"net/http"

	"github.com/hiromaily/go-graphql-server/pkg/server/handler"
	"github.com/hiromaily/go-graphql-server/pkg/server/httpmethod"
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
	method httpmethod.HTTPMethod,
	port int,
) Server {
	return newServer(
		logger,
		schema,
		method,
		port,
	)
}

// server object
type server struct {
	logger *zap.Logger
	schema graphql.Schema
	method httpmethod.HTTPMethod
	port   int
}

// newServer returns server object
func newServer(
	logger *zap.Logger,
	schema graphql.Schema,
	method httpmethod.HTTPMethod,
	port int,
) *server {
	return &server{
		logger: logger,
		schema: schema,
		method: method,
		port:   port,
	}
}

// Start starts server
func (s *server) Start() error {
	if err := handler.Initialize(s.schema, s.method); err != nil {
		return err
	}

	s.logger.Info("server is running", zap.Int("port", s.port))
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
	http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)

	return nil
}

//func (s *server) StartTest() *httptest.Server {
//	httptest.NewServer()
//}

// Close closes dependencies
func (s *server) Close() {
	// TODO: DB close
}
