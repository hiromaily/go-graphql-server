package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/hiromaily/go-graphql-server/pkg/model/company"
	"github.com/hiromaily/go-graphql-server/pkg/model/country"
	"github.com/hiromaily/go-graphql-server/pkg/model/user"
	"github.com/hiromaily/go-graphql-server/pkg/model/workhistory"
)

func newQueryType(
	userResolver user.UserFieldResolver,
	companyResolver company.CompanyFieldResolver,
	countryResolver country.CountryFieldResolver,
	workHistoryResolver workhistory.WorkHistoryFieldResolver,
) *graphql.Object {
	/*
	   Create Query object type with fields "user" has type [userType] by using GraphQLObjectTypeConfig:
	       - Name: name of object type
	       - Fields: a map of fields by using GraphQLFields
	   Setup type of field use GraphQLFieldConfig to define:
	       - Type: type of field
	       - Args: arguments to query with current field
	       - Resolve: function to query data using params from [Args] and return value with current type
	*/
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				/*
				   curl -g 'http://localhost:8080/graphql?query={user(id:"1"){name,age,country}}'
				*/
				"user": &graphql.Field{
					Type: userType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: userResolver.GetByID,
				},
				/*
				   curl -g 'http://localhost:8080/graphql?query={userList{id,name}}'
				*/
				"userList": &graphql.Field{
					Type:        graphql.NewList(userType),
					Description: "List of user",
					Resolve:     userResolver.List,
				},
				/*
				   curl -g 'http://localhost:8080/graphql?query={company(id:"1"){id,name,country}}'
				*/
				"company": &graphql.Field{
					Type: companyType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: companyResolver.GetByID,
				},
				/*
				   curl -g 'http://localhost:8080/graphql?query={companyList{id,name}}'
				*/
				"companyList": &graphql.Field{
					Type:        graphql.NewList(companyType),
					Description: "List of company",
					Resolve:     companyResolver.List,
				},
				/*
				   curl -g 'http://localhost:8080/graphql?query={country(id:"1"){id,name,code}}'
				*/
				"country": &graphql.Field{
					Type: countryType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: countryResolver.GetByID,
				},
				/*
				   curl -g 'http://localhost:8080/graphql?query={countryList{id,name}}'
				*/
				"countryList": &graphql.Field{
					Type:        graphql.NewList(countryType),
					Description: "List of country",
					Resolve:     countryResolver.List,
				},
				/*
				   curl -g 'http://localhost:8080/graphql?query={workHistory(id:"1"){id,company,title}}'
				*/
				"workHistory": &graphql.Field{
					Type: workHistoryType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: workHistoryResolver.GetByID,
				},
				/*
				   curl -g 'http://localhost:8080/graphql?query={userWorkHistory(user_id:"1"){id,company,title}}'
				*/
				"userWorkHistory": &graphql.Field{
					Type: workHistoryType,
					Args: graphql.FieldConfigArgument{
						"user_id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: workHistoryResolver.GetByUserID,
				},
				/*
				   curl -g 'http://localhost:8080/graphql?query={workHistoryList(){id,company,title}}'
				*/
				"workHistoryList": &graphql.Field{
					Type:        graphql.NewList(workHistoryType),
					Description: "List of work history",
					Resolve:     workHistoryResolver.List,
				},
			},
		},
	)

	return queryType
}
