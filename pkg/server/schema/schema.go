package schema

import (
	"github.com/graphql-go/graphql"

	"github.com/hiromaily/go-graphql-server/pkg/user"
)

/*
   Create User object type with fields "id" and "name" by using GraphQLObjectTypeConfig:
       - Name: name of object type
       - Fields: a map of fields by using GraphQLFields
   Setup type of field use GraphQLFieldConfig
*/
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func newQueryType(userResolver user.UserFieldResolver) *graphql.Object {
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
				   curl -g 'http://localhost:8080/graphql?query={user(id:"1"){name}}'
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
					Resolve:     userResolver.UserList,
				},
			},
		})

	return queryType
}

func NewSchema(userResolver user.UserFieldResolver) graphql.Schema {
	// schema
	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: newQueryType(userResolver),
		},
	)
	return schema
}
