package repository

import (
	"strconv"

	"github.com/pkg/errors"

	"github.com/hiromaily/go-graphql-server/pkg/files"
	"github.com/hiromaily/go-graphql-server/pkg/model/company"
)

type companyMap struct {
	repo map[string]company.CompanyType
	list []*company.CompanyType
}

// NewCompanyMapRepo returns Company interface
func NewCompanyMapRepo() (company.Company, error) {
	var data map[string]company.CompanyType
	err := files.ImportJSONFile("./assets/company.json", &data)
	if err != nil {
		return nil, err
	}
	return &companyMap{
		repo: data,
	}, nil
}

func (c *companyMap) updateList() {
	ctList := make([]*company.CompanyType, 0, len(c.repo))
	for _, val := range c.repo {
		val := val
		ctList = append(ctList, &val)
	}
	c.list = ctList
}

// Fetch returns company by id
func (c *companyMap) Fetch(id string) (*company.CompanyType, error) {
	if v, ok := c.repo[id]; ok {
		return &v, nil
	}
	return nil, errors.New("company is not found")
}

// FetchByName returns company by name
func (c *companyMap) FetchByName(name string) (*company.CompanyType, error) {
	return nil, errors.New("not implemented")
}

// FetchAll returns all companies
func (c *companyMap) FetchAll() ([]*company.CompanyType, error) {
	if len(c.list) == 0 {
		c.updateList()
	}
	return c.list, nil
}

func (c *companyMap) Insert(ct *company.CompanyType) error {
	id := strconv.Itoa(ct.ID)
	if _, ok := c.repo[id]; ok {
		return errors.Errorf("id[%d] is already existing", ct.ID)
	}
	c.repo[id] = *ct
	c.list = append(c.list, ct)

	return nil
}

func (c *companyMap) Update(ct *company.CompanyType) error {
	id := strconv.Itoa(ct.ID)
	if _, ok := c.repo[id]; !ok {
		return errors.Errorf("id[%d] is not found", ct.ID)
	}
	updated, err := c.Fetch(id)
	if err != nil {
		return err
	}
	if ct.Name == "" {
		ct.Name = updated.Name
	}
	c.repo[id] = *ct
	c.updateList()

	return nil
}

func (c *companyMap) Delete(id string) error {
	delete(c.repo, id)
	c.updateList()
	return nil
}
