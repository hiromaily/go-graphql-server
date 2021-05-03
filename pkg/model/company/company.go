package company

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Company for fetching data interface
// - implementation is in repository
type Company interface {
	Fetch(id string) (*CompanyType, error)
	FetchAll() ([]*CompanyType, error)
	Insert(ct *CompanyType) error
	Update(ct *CompanyType) error
	Delete(id string) error
}

// CompanyType is type of company
type CompanyType struct {
	ID      int    `json:"id" boil:"id"`
	Name    string `json:"name" boil:"name"`
	Country string `json:"country" boil:"country"`
}

// CompanyFieldResolver for resolver of schema interface
type CompanyFieldResolver interface {
	GetByID(p graphql.ResolveParams) (interface{}, error)
	List(p graphql.ResolveParams) (interface{}, error)
	Create(p graphql.ResolveParams) (interface{}, error)
	Update(p graphql.ResolveParams) (interface{}, error)
	Delete(p graphql.ResolveParams) (interface{}, error)
}

type companyFieldResolver struct {
	logger      *zap.Logger
	companyRepo Company
}

// NewCompanyFieldResolve returns CompanyFieldResolver interface
func NewCompanyFieldResolve(
	logger *zap.Logger,
	companyRepo Company,
) CompanyFieldResolver {
	return &companyFieldResolver{
		logger:      logger,
		companyRepo: companyRepo,
	}
}

// GetByID gets company by ID
func (c *companyFieldResolver) GetByID(p graphql.ResolveParams) (interface{}, error) {
	idQuery, isOK := p.Args["id"].(string)
	if isOK {
		return c.companyRepo.Fetch(idQuery)
	}
	return nil, errors.New("not found")
}

// List returns all companies
func (c *companyFieldResolver) List(_ graphql.ResolveParams) (interface{}, error) {
	return c.companyRepo.FetchAll()
}

// Create creates new user by parameters
func (c *companyFieldResolver) Create(p graphql.ResolveParams) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	newCompany := &CompanyType{
		ID:      rand.Intn(100000), // TODO: get maximum ID from list
		Name:    p.Args["name"].(string),
		Country: p.Args["country"].(string),
	}
	// insert to repository
	err := c.companyRepo.Insert(newCompany)
	if err != nil {
		return nil, err
	}
	return newCompany, nil
}

func (c *companyFieldResolver) Update(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	updated := CompanyType{
		ID: intID,
	}

	if name, ok := p.Args["name"].(string); ok {
		updated.Name = name
	}
	if country, ok := p.Args["country"].(string); ok {
		updated.Country = country
	}
	if err := c.companyRepo.Update(&updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (c *companyFieldResolver) Delete(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	deleted, err := c.companyRepo.Fetch(id)
	if err != nil {
		return nil, err
	}
	c.companyRepo.Delete(id)

	return deleted, nil
}
