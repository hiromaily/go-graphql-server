package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/hiromaily/go-graphql-server/pkg/model/company"
	"github.com/hiromaily/go-graphql-server/pkg/model/country"
	"github.com/hiromaily/go-graphql-server/pkg/model/user"
	"github.com/hiromaily/go-graphql-server/pkg/model/workhistory"
)

// NewSchema returns graphql.Schema
func NewSchema(
	userResolver user.UserFieldResolver,
	companyResolver company.CompanyFieldResolver,
	countryResolver country.CountryFieldResolver,
	workHistoryResolver workhistory.WorkHistoryFieldResolver,
) graphql.Schema {
	// schema
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    newQueryType(userResolver, companyResolver, countryResolver, workHistoryResolver),
			Mutation: newMutationType(userResolver, companyResolver),
			// Subscription: // TODO: implementation
		},
	)
	return schema
}
