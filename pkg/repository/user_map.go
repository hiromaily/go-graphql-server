package repository

import (
	"strconv"

	"github.com/pkg/errors"

	"github.com/hiromaily/go-graphql-server/pkg/files"
	"github.com/hiromaily/go-graphql-server/pkg/user"
)

type userMap struct {
	repo map[string]user.UserType
	list []*user.UserType
}

// NewUserMapRepo returns User interface
func NewUserMapRepo() (user.User, error) {
	var data map[string]user.UserType
	err := files.ImportJSONFile("./assets/user.json", &data)
	if err != nil {
		return nil, err
	}
	return &userMap{
		repo: data,
	}, nil
}

func (u *userMap) updateList() {
	utList := make([]*user.UserType, 0, len(u.repo))
	for _, val := range u.repo {
		val := val
		utList = append(utList, &val)
	}
	u.list = utList
}

// Fetch returns user by id
func (u *userMap) Fetch(id string) (*user.UserType, error) {
	if v, ok := u.repo[id]; ok {
		return &v, nil
	}
	return nil, errors.New("user is not found")
}

// FetchAll returns all users
func (u *userMap) FetchAll() ([]*user.UserType, error) {
	if len(u.list) == 0 {
		u.updateList()
	}
	return u.list, nil
}

func (u *userMap) Insert(ut *user.UserType) error {
	id := strconv.Itoa(ut.ID)
	if _, ok := u.repo[id]; ok {
		return errors.Errorf("id[%d] is already existing", ut.ID)
	}
	u.repo[id] = *ut
	u.list = append(u.list, ut)

	return nil
}

func (u *userMap) Update(ut *user.UserType) error {
	id := strconv.Itoa(ut.ID)
	if _, ok := u.repo[id]; !ok {
		return errors.Errorf("id[%d] is not found", ut.ID)
	}
	u.repo[id] = *ut
	u.updateList()

	return nil
}

func (u *userMap) Delete(id string) {
	delete(u.repo, id)
	u.updateList()
}
