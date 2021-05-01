package repository

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"

	"github.com/hiromaily/go-graphql-server/pkg/country"
	models "github.com/hiromaily/go-graphql-server/pkg/model/rdb"
)

type countryDB struct {
	dbConn    *sql.DB
	tableName string
	logger    *zap.Logger
}

// NewCountryDBRepo returns Country interface
func NewCountryDBRepo(dbConn *sql.DB, logger *zap.Logger) country.Country {
	return &countryDB{
		dbConn:    dbConn,
		tableName: "m_country",
		logger:    logger,
	}
}

// Fetch returns country by id
func (c *countryDB) Fetch(id string) (*country.CountryType, error) {
	ctx := context.Background()

	var country *country.CountryType
	err := models.MCountries(
		qm.Select("id, country_code, name"),
		qm.Where("id=?", id),
	).Bind(ctx, c.dbConn, country)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.MCountries().Bind()")
	}

	return country, nil
}

// FetchByName returns country by name
func (c *countryDB) FetchByName(name string) (*country.CountryType, error) {
	ctx := context.Background()

	var country *country.CountryType
	err := models.MCountries(
		qm.Select("id, country_code, name"),
		qm.Where("name=?", name),
	).Bind(ctx, c.dbConn, country)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.MCountries().Bind()")
	}

	return country, nil
}

// FetchAll returns all countries
func (c *countryDB) FetchAll() ([]*country.CountryType, error) {
	ctx := context.Background()

	var countries []*country.CountryType
	// sql := "SELECT id FROM t_users WHERE delete_flg=?"
	err := models.MCountries(
		qm.Select("id, country_code, name"),
	).Bind(ctx, c.dbConn, &countries)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.MCountries().Bind()")
	}
	return countries, nil
}
