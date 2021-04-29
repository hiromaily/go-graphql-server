package user

import (
	"github.com/graphql-go/graphql"
	"go.uber.org/zap"
)

type User interface {
	Fetch(id string) UserType
}

type UserType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserFieldResolver interface {
	GetByID(p graphql.ResolveParams) (interface{}, error)
}

type userFieldResolver struct {
	logger  *zap.Logger
	fetcher User
}

func NewUserFieldResolveFn(
	logger *zap.Logger,
	userFetcher User,
) UserFieldResolver {
	return &userFieldResolver{
		logger:  logger,
		fetcher: userFetcher,
	}
}

func (u *userFieldResolver) GetByID(p graphql.ResolveParams) (interface{}, error) {
	idQuery, isOK := p.Args["id"].(string)
	if isOK {
		return u.fetcher.Fetch(idQuery), nil
	}
	return nil, nil
}
