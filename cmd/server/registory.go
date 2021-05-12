package main

import (
	"database/sql"


	"github.com/graphql-go/graphql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"

	"github.com/hiromaily/go-graphql-server/pkg/config"
	"github.com/hiromaily/go-graphql-server/pkg/db/mysql"
	"github.com/hiromaily/go-graphql-server/pkg/logger"
	"github.com/hiromaily/go-graphql-server/pkg/model/company"
	"github.com/hiromaily/go-graphql-server/pkg/model/country"
	"github.com/hiromaily/go-graphql-server/pkg/model/user"
	"github.com/hiromaily/go-graphql-server/pkg/model/workhistory"
	"github.com/hiromaily/go-graphql-server/pkg/repository"
	"github.com/hiromaily/go-graphql-server/pkg/schema"
	"github.com/hiromaily/go-graphql-server/pkg/server"
)

// Registry interface
type Registry interface {
	NewServer() server.Server
}

type registry struct {
	conf            *config.Root
	logger          *zap.Logger
	mysqlClient     *sql.DB
	userRepo        user.User
	companyRepo     company.Company
	countryRepo     country.Country
	workHistoryRepo workhistory.WorkHistory
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
		r.newCompanyFieldResolver(),
		r.newCountryFieldResolver(),
		r.newWorkHistoryResolver(),
	)
}

func (r *registry) newUserFieldResolver() user.UserFieldResolver {
	return user.NewUserFieldResolve(
		r.newLogger(),
		r.newUserRepo(),
	)
}

func (r *registry) newCompanyFieldResolver() company.CompanyFieldResolver {
	return company.NewCompanyFieldResolve(
		r.newLogger(),
		r.newComanyRepo(),
	)
}

func (r *registry) newCountryFieldResolver() country.CountryFieldResolver {
	return country.NewCountryFieldResolve(
		r.newLogger(),
		r.newCountryRepo(),
	)
}

func (r *registry) newWorkHistoryResolver() workhistory.WorkHistoryFieldResolver {
	return workhistory.NewWorkHistoryFieldResolve(
		r.newLogger(),
		r.newWorkHistoryRepo(),
	)
}

func (r *registry) newUserRepo() user.User {
	if r.userRepo == nil {
		if r.conf.MySQL.IsEnabled {
			// using DB
			r.userRepo = repository.NewUserDBRepo(
				r.newMySQLClient(),
				r.newLogger(),
				r.newCountryRepo(),
			)
		} else {
			// map pattern
			repo, err := repository.NewUserMapRepo()
			if err != nil {
				panic(err)
			}
			r.userRepo = repo
		}
	}
	return r.userRepo
}

func (r *registry) newComanyRepo() company.Company {
	if r.companyRepo == nil {
		if r.conf.MySQL.IsEnabled {
			// using DB
			r.companyRepo = repository.NewCompanyDBRepo(
				r.newMySQLClient(),
				r.newLogger(),
				r.newCountryRepo(),
			)
		} else {
			// map pattern
			repo, err := repository.NewCompanyMapRepo()
			if err != nil {
				panic(err)
			}
			r.companyRepo = repo
		}
	}
	return r.companyRepo
}

func (r *registry) newCountryRepo() country.Country {
	if r.countryRepo == nil {
		if r.conf.MySQL.IsEnabled {
			// using DB
			r.countryRepo = repository.NewCountryDBRepo(
				r.newMySQLClient(),
				r.newLogger(),
			)
		} else {
			// map pattern
			repo, err := repository.NewCountryMapRepo()
			if err != nil {
				panic(err)
			}
			r.countryRepo = repo
		}
	}
	return r.countryRepo
}

func (r *registry) newWorkHistoryRepo() workhistory.WorkHistory {
	if r.workHistoryRepo == nil {
		if r.conf.MySQL.IsEnabled {
			// using DB
			r.workHistoryRepo = repository.NewWorkHistoryDBRepo(
				r.newMySQLClient(),
				r.newLogger(),
				r.newComanyRepo(),
			)
		} else {
			// map pattern
			repo, err := repository.NewWorkHistoryMapRepo()
			if err != nil {
				panic(err)
			}
			r.workHistoryRepo = repo
		}
	}
	return r.workHistoryRepo
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
