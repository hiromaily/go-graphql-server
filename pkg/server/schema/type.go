package schema

import (
	"github.com/graphql-go/graphql"
)

// Scalar Type
// - graphql.Int
// - graphql.Float
// - graphql.String
// - graphql.Boolean
// - graphql.ID

// String! => graphql.NewNonNull(graphql.String)

// Custom scalar type can validate value
// e.g. DateTime type, see example

// Enum
//var enumType = graphql.NewEnum(graphql.EnumConfig{
//	Name: "Enum",
//	Values: graphql.EnumValueConfigMap{
//		"foo": &graphql.EnumValueConfig{},
//	},
//})

// List
// - graphql.NewList(graphql.Int)
//    => list may be null, content may be null as well
// - graphql.NewList(graphql.NewNonNull(graphql.Int))
//    => list may be null, but content is not null
// - graphql.NewNonNull(graphql.NewList(graphql.Int))
//    => list is not null, but content may be null
// - graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(graphql.Int)))
//    => list is not null, content is not null
//
// empty list `[]` is recommended to return when no value


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
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
			"country": &graphql.Field{
				Type: graphql.String,
			},
			"resume": &graphql.Field{
				Type: workHistoryType,
			},
		},
	},
)

var workHistoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "WorkHisotry",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"company": &graphql.Field{
				Type: companyType,
			},
			"started_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"ended_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

var companyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Company",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"country": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

