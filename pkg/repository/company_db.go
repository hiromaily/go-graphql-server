package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"

	"github.com/hiromaily/go-graphql-server/pkg/model/company"
	"github.com/hiromaily/go-graphql-server/pkg/model/country"
	models "github.com/hiromaily/go-graphql-server/pkg/model/rdb"
)

type companyDB struct {
	dbConn    *sql.DB
	tableName string
	logger    *zap.Logger
	country   country.Country
}

// NewCompanyDBRepo returns Company interface
func NewCompanyDBRepo(dbConn *sql.DB, logger *zap.Logger, country country.Country) company.Company {
	return &companyDB{
		dbConn:    dbConn,
		tableName: "t_company",
		logger:    logger,
		country:   country,
	}
}

// Fetch returns company by id
func (c *companyDB) Fetch(id string) (*company.CompanyType, error) {
	ctx := context.Background()

	var company company.CompanyType
	err := models.TCompanies(
		qm.Select("t_company.id, t_company.name, cty.name as country"),
		qm.LeftOuterJoin("m_country as cty on t_company.country_id = cty.id"),
		qm.Where("t_company.id=?", id),
	).Bind(ctx, c.dbConn, &company)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TCompanies().Bind() in Fetch()")
	}

	return &company, nil
}

// FetchAll returns all companies
func (c *companyDB) FetchAll() ([]*company.CompanyType, error) {
	ctx := context.Background()

	var companies []*company.CompanyType
	err := models.TCompanies(
		qm.Select("t_company.id, t_company.name, cty.name as country"),
		qm.LeftOuterJoin("m_country as cty on t_company.country_id = cty.id"),
	).Bind(ctx, c.dbConn, &companies)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TCompanies().Bind() in FetchAll()")
	}
	return companies, nil
}

func (c *companyDB) Insert(ct *company.CompanyType) error {
	// get country
	countryType, err := c.country.FetchByName(ct.Country)
	if err != nil {
		return err
	}

	item := &models.TCompany{
		Name:      ct.Name,
		CountryID: uint8(countryType.ID),
	}

	ctx := context.Background()

	if err := item.Insert(ctx, c.dbConn, boil.Infer()); err != nil {
		return errors.Wrap(err, "failed to call company.Insert()")
	}
	// TODO: return latest ID
	return nil
}

func (c *companyDB) Update(ct *company.CompanyType) error {
	ctx := context.Background()

	// Set updating columns
	updCols := map[string]interface{}{}
	if ct.Name != "" {
		updCols[models.TCompanyColumns.Name] = ct.Name
	}
	if ct.Country != "" {
		cty, err := c.country.FetchByName(ct.Country)
		if err != nil {
			return err
		}
		updCols[models.TCompanyColumns.CountryID] = cty.ID
	}
	updCols[models.TCompanyColumns.UpdatedAt] = null.TimeFrom(time.Now().UTC())

	_, err := models.TCompanies(
		qm.Where("id=?", ct.ID),
	).UpdateAll(ctx, c.dbConn, updCols)

	return err
}

func (c *companyDB) Delete(id string) error {
	ctx := context.Background()

	_, err := models.TCompanies(
		qm.Where("t_company.id=?", id),
	).DeleteAll(ctx, c.dbConn)
	return err
}
