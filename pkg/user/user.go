package user

import (
	"github.com/graphql-go/graphql"
	"go.uber.org/zap"
)

// User for fetching data interface
type User interface {
	Fetch(id string) UserType
	FetchAll() []UserType
}

// UserType is type of user
type UserType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UserFieldResolver for resolver of schema interface
type UserFieldResolver interface {
	GetByID(p graphql.ResolveParams) (interface{}, error)
	UserList(p graphql.ResolveParams) (interface{}, error)
}

type userFieldResolver struct {
	logger  *zap.Logger
	fetcher User
}

// NewUserFieldResolve returns UserFieldResolver interface
func NewUserFieldResolve(
	logger *zap.Logger,
	userFetcher User,
) UserFieldResolver {
	return &userFieldResolver{
		logger:  logger,
		fetcher: userFetcher,
	}
}

// GetByID gets user by ID
func (u *userFieldResolver) GetByID(p graphql.ResolveParams) (interface{}, error) {
	idQuery, isOK := p.Args["id"].(string)
	if isOK {
		return u.fetcher.Fetch(idQuery), nil
	}
	return nil, nil
}

// UserList returns all users
func (u *userFieldResolver) UserList(_ graphql.ResolveParams) (interface{}, error) {
	return u.fetcher.FetchAll(), nil
}
