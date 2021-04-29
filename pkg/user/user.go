package user

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"go.uber.org/zap"
)

// User for fetching data interface
type User interface {
	Fetch(id string) UserType
	FetchAll() []UserType
	Insert(ut UserType) error
	Update(ut UserType) error
	Delete(id string)
}

// UserType is type of user
type UserType struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

// UserFieldResolver for resolver of schema interface
type UserFieldResolver interface {
	GetByID(p graphql.ResolveParams) (interface{}, error)
	List(p graphql.ResolveParams) (interface{}, error)
	Create(p graphql.ResolveParams) (interface{}, error)
	Update(p graphql.ResolveParams) (interface{}, error)
	Delete(p graphql.ResolveParams) (interface{}, error)
}

type userFieldResolver struct {
	logger   *zap.Logger
	userRepo User
}

// NewUserFieldResolve returns UserFieldResolver interface
func NewUserFieldResolve(
	logger *zap.Logger,
	userRepo User,
) UserFieldResolver {
	return &userFieldResolver{
		logger:   logger,
		userRepo: userRepo,
	}
}

// GetByID gets user by ID
func (u *userFieldResolver) GetByID(p graphql.ResolveParams) (interface{}, error) {
	idQuery, isOK := p.Args["id"].(string)
	if isOK {
		return u.userRepo.Fetch(idQuery), nil
	}
	return nil, nil
}

// UserList returns all users
func (u *userFieldResolver) List(_ graphql.ResolveParams) (interface{}, error) {
	return u.userRepo.FetchAll(), nil
}

// Create creates new user by parameters
func (u *userFieldResolver) Create(p graphql.ResolveParams) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	newUser := UserType{
		ID:      strconv.Itoa(rand.Intn(100000)), // TODO: get maximum ID
		Name:    p.Args["name"].(string),
		Age:     p.Args["age"].(int),
		Country: p.Args["country"].(string),
	}
	// insert to repository
	err := u.userRepo.Insert(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *userFieldResolver) Update(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	updated := u.userRepo.Fetch(id)

	if name, ok := p.Args["name"].(string); ok {
		updated.Name = name
	}
	if age, ok := p.Args["age"].(int); ok {
		updated.Age = age
	}
	if country, ok := p.Args["country"].(string); ok {
		updated.Country = country
	}
	if err := u.userRepo.Update(updated); err != nil {
		return nil, err
	}

	return updated, nil
}

func (u *userFieldResolver) Delete(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	deleted := u.userRepo.Fetch(id)
	u.userRepo.Delete(id)

	return deleted, nil
}
