package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/hiromaily/go-graphql-server/pkg/model/company"
	"github.com/hiromaily/go-graphql-server/pkg/model/user"
)

func newMutationType(
	userResolver user.UserFieldResolver,
	companyResolver company.CompanyFieldResolver,
) *graphql.Object {
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			/*
			   curl -g 'http://localhost:8080/graphql?query=mutation+_{createUser(name:"Tom",age:15,country:"Japan"){id,name,age,country}}'
			*/
			"createUser": &graphql.Field{
				Type:        userType,
				Description: "Create new user",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"country": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: userResolver.Create,
			},
			/*
				curl -g 'http://localhost:8080/graphql?query=mutation+_{updateUser(id:"1",name:"Dummy",age:99,country:"Japan"){id,name,age,country}}'
			*/
			"updateUser": &graphql.Field{
				Type:        userType,
				Description: "Update user by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"country": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: userResolver.Update,
			},
			/*
				curl -g 'http://localhost:8080/graphql?query=mutation+_{deleteUser(id:"2"){id,name,age,country}}'
			*/
			"deleteUser": &graphql.Field{
				Type:        userType,
				Description: "Delete user by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: userResolver.Delete,
			},
			/*
			   curl -g 'http://localhost:8080/graphql?query=mutation+_{createCompany(name:"TechTech",country:"Japan"){id,name,country}}'
			*/
			"createCompany": &graphql.Field{
				Type:        companyType,
				Description: "Create new company",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"country": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: companyResolver.Create,
			},
			/*
				curl -g 'http://localhost:8080/graphql?query=mutation+_{updateCompany(id:"1",name:"TechTechTech"){id,name,country}}'
			*/
			"updateCompany": &graphql.Field{
				Type:        companyType,
				Description: "Update company by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"country": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: companyResolver.Update,
			},
			/*
				curl -g 'http://localhost:8080/graphql?query=mutation+_{deleteCompany(id:"2"){id,name,country}}'
			*/
			"deleteCompany": &graphql.Field{
				Type:        userType,
				Description: "Delete company by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: companyResolver.Delete,
			},
		},
	})

	return mutationType
}
