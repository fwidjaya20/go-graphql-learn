package types

import "github.com/graphql-go/graphql"

var User = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "User",
		Fields:      graphql.Fields{
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"full_name": &graphql.Field{
				Type: graphql.String,
			},
			"identity_card": &graphql.Field{
				Type: graphql.String,
			},
			"role_id": &graphql.Field{
				Type: graphql.Int,
			},
			"division_id": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)