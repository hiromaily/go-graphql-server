package country

import (
	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Country for fetching data interface
type Country interface {
	Fetch(id string) (*CountryType, error)
	FetchByName(name string) (*CountryType, error)
	FetchAll() ([]*CountryType, error)
}

// CountryType is type of user
type CountryType struct {
	ID   int    `json:"id" boil:"id"`
	Code string `json:"country_code" boil:"country_code"`
	Name string `json:"name" boil:"name"`
}

// CountryFieldResolver for resolver of schema interface
type CountryFieldResolver interface {
	GetByID(p graphql.ResolveParams) (interface{}, error)
	List(p graphql.ResolveParams) (interface{}, error)
}

type countryFieldResolver struct {
	logger      *zap.Logger
	countryRepo Country
}

// NewCountryFieldResolve returns CountryFieldResolver interface
func NewCountryFieldResolve(
	logger *zap.Logger,
	countryRepo Country,
) CountryFieldResolver {
	return &countryFieldResolver{
		logger:      logger,
		countryRepo: countryRepo,
	}
}

// GetByID gets country by ID
func (c *countryFieldResolver) GetByID(p graphql.ResolveParams) (interface{}, error) {
	idQuery, isOK := p.Args["id"].(string)
	if isOK {
		return c.countryRepo.Fetch(idQuery)
	}
	return nil, errors.New("not found")
}

// List returns all countries
func (c *countryFieldResolver) List(_ graphql.ResolveParams) (interface{}, error) {
	return c.countryRepo.FetchAll()
}
