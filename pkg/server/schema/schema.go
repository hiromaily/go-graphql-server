package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/hiromaily/go-graphql-server/pkg/model/country"
	"github.com/hiromaily/go-graphql-server/pkg/model/user"
)

// NewSchema returns graphql.Schema
func NewSchema(
	userResolver user.UserFieldResolver,
	countryResolver country.CountryFieldResolver) graphql.Schema {
	// schema
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    newQueryType(userResolver, countryResolver),
			Mutation: newMutationType(userResolver),
		},
	)
	return schema
}
