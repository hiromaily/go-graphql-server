package repository

import (
	"github.com/pkg/errors"

	"github.com/hiromaily/go-graphql-server/pkg/country"
	"github.com/hiromaily/go-graphql-server/pkg/files"
)

type countryMap struct {
	repo map[string]country.CountryType
	list []*country.CountryType
}

// NewCountryMapRepo returns Country interface
func NewCountryMapRepo() (country.Country, error) {
	var data map[string]country.CountryType
	err := files.ImportJSONFile("./assets/country.json", &data)
	if err != nil {
		return nil, err
	}
	return &countryMap{
		repo: data,
	}, nil
}

func (c *countryMap) updateList() {
	ctList := make([]*country.CountryType, 0, len(c.repo))
	for _, val := range c.repo {
		val := val
		ctList = append(ctList, &val)
	}
	c.list = ctList
}

// Fetch returns user by id
func (c *countryMap) Fetch(id string) (*country.CountryType, error) {
	if v, ok := c.repo[id]; ok {
		return &v, nil
	}
	return nil, errors.New("user is not found")
}

// FetchByName returns user by name
func (c *countryMap) FetchByName(name string) (*country.CountryType, error) {
	return nil, errors.New("not implemented")
}

// FetchAll returns all users
func (c *countryMap) FetchAll() ([]*country.CountryType, error) {
	if len(c.list) == 0 {
		c.updateList()
	}
	return c.list, nil
}
