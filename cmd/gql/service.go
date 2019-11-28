package gql

import (
	"github.com/fwidjaya20/go-graphql-learn/cmd/containers"
	"github.com/fwidjaya20/go-graphql-learn/internal/gql"
	gqlHttpTransport "github.com/fwidjaya20/go-graphql-learn/internal/gql/transports/http"
	"github.com/graphql-go/graphql"
	"log"
)

func InitGQL(resolverContainter containers.ResolverContainer) gqlHttpTransport.Server {
	rootQuery := gql.NewRoot(resolverContainter)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:        rootQuery.Query,
			Mutation:     nil,
			Subscription: nil,
			Types:        nil,
			Directives:   nil,
			Extensions:   nil,
		},
	)

	if nil != err {
		log.Println("ERROR_CreateGQLSchema", err)
	}

	return gqlHttpTransport.Server{
		GqlSchema: &sc,
	}
}