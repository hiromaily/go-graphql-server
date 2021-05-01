package main

import (
	"database/sql"
	"github.com/graphql-go/graphql"
	"github.com/hiromaily/go-graphql-server/pkg/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"

	"github.com/hiromaily/go-graphql-server/pkg/config"
	"github.com/hiromaily/go-graphql-server/pkg/logger"
	"github.com/hiromaily/go-graphql-server/pkg/repository"
	"github.com/hiromaily/go-graphql-server/pkg/server"
	"github.com/hiromaily/go-graphql-server/pkg/server/schema"
	"github.com/hiromaily/go-graphql-server/pkg/user"
)

// Registry interface
type Registry interface {
	NewServer() server.Server
}

type registry struct {
	conf     *config.Root
	logger   *zap.Logger
	userRepo user.User
	mysqlClient *sql.DB
}

// NewRegistry is to register regstry interface
func NewRegistry(conf *config.Root) Registry {
	return &registry{conf: conf}
}

// NewServer registers for Server interface
func (r *registry) NewServer() server.Server {
	return server.NewServer(
		r.newLogger(),
		r.newSchema(),
		r.conf.Server.HTTPMethod,
		r.conf.Server.Port,
	)
}

func (r *registry) newLogger() *zap.Logger {
	if r.logger == nil {
		r.logger = logger.NewZapLogger(r.conf.Logger)
	}
	return r.logger
}

func (r *registry) newSchema() graphql.Schema {
	return schema.NewSchema(
		r.newUserFieldResolver(),
	)
}

func (r *registry) newUserFieldResolver() user.UserFieldResolver {
	return user.NewUserFieldResolve(
		r.newLogger(),
		r.newUserRepo(),
	)
}

func (r *registry) newUserRepo() user.User {
	if r.userRepo == nil {
		// using DB
		repo := repository.NewUserDBRepo(
			r.newMySQLClient(),
			r.newLogger(),
		)
		// map pattern
		//repo, err := repository.NewUserMapRepo()
		//if err != nil {
		//	panic(err)
		//}
		r.userRepo = repo
	}
	return r.userRepo
}

func (r *registry) newMySQLClient() *sql.DB {
	if r.mysqlClient == nil {
		dbConn, err := mysql.NewMySQL(r.conf.MySQL)
		if err != nil {
			panic(err)
		}
		r.mysqlClient = dbConn
		if r.conf.MySQL.IsDebugLog {
			boil.DebugMode = true
		}
	}
	return r.mysqlClient
}
