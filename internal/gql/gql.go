package gql

import (
	"fmt"
	"github.com/fwidjaya20/go-graphql-learn/cmd/containers"
	"github.com/fwidjaya20/go-graphql-learn/internal/gql/types"
	"github.com/graphql-go/graphql"
)

type Root struct {
	Query *graphql.Object
}

func NewRoot(resolverContainer containers.ResolverContainer) *Root {
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						Type: graphql.NewList(types.User),
						Args: graphql.FieldConfigArgument{
							"full_name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolverContainer.UserResolver.Resolve,
					},
				},
			},
		),
	}

	return &root
}

func ExecureQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("Unexpected errors inside ExecuteQuery: %v", result.Errors)
	}

	return result
}