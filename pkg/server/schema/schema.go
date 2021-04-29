package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/hiromaily/go-graphql-server/pkg/user"
)

// NewSchema returns graphql.Schema
func NewSchema(userResolver user.UserFieldResolver) graphql.Schema {
	// schema
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    newQueryType(userResolver),
			Mutation: newMutationType(userResolver),
		},
	)
	return schema
}
