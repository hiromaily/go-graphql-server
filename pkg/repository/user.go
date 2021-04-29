package repository

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"

	"github.com/hiromaily/go-graphql-server/pkg/user"
)

type userMap struct {
	repo map[string]user.UserType
	list []user.UserType
}

// NewUserMapRepo returns User interface
func NewUserMapRepo() (user.User, error) {
	data, err := importJSONFile("./assets/user.json")
	if err != nil {
		return nil, err
	}
	return &userMap{
		repo: data,
	}, nil
}

// importJSONFile imports json file to map
func importJSONFile(fileName string) (map[string]user.UserType, error) {
	var data map[string]user.UserType
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	// err = json.Unmarshal(content, &data)
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *userMap) updateList() {
	utList := make([]user.UserType, 0, len(u.repo))
	for _, val := range u.repo {
		utList = append(utList, val)
	}
	u.list = utList
}

// Fetch returns user by id
func (u *userMap) Fetch(id string) user.UserType {
	return u.repo[id]
}

// FetchAll returns all users
func (u *userMap) FetchAll() []user.UserType {
	if len(u.list) == 0 {
		u.updateList()
	}
	return u.list
}

func (u *userMap) Insert(ut user.UserType) error {
	if _, ok := u.repo[ut.ID]; ok {
		return errors.Errorf("id[%s] is already existing", ut.ID)
	}
	u.repo[ut.ID] = ut
	u.list = append(u.list, ut)

	return nil
}

func (u *userMap) Update(ut user.UserType) error {
	if _, ok := u.repo[ut.ID]; !ok {
		return errors.Errorf("id[%s] is not found", ut.ID)
	}
	u.repo[ut.ID] = ut
	u.updateList()

	return nil
}

func (u *userMap) Delete(id string) {
	delete(u.repo, id)
	u.updateList()
}
