package main

import (
	"github.com/fwidjaya20/go-graphql-learn/cmd/containers"
	"github.com/fwidjaya20/go-graphql-learn/cmd/gql"
	"github.com/fwidjaya20/go-graphql-learn/config"
	"github.com/go-chi/chi"
	"github.com/oklog/oklog/pkg/group"
	"log"
	"net/http"
)

func main() {
	var g group.Group

	resolverContainer := containers.NewResolverContainer()

	initHTTP(&g, resolverContainer)

	log.Fatal("exit", g.Run())
}

func initHTTP(g *group.Group, resolverContainer containers.ResolverContainer) {
	HttpAddr := config.GetEnv(config.HTTP_ADDR)

	var router *chi.Mux
	router = chi.NewRouter()

	s := http.Server{
		Addr:    HttpAddr,
		Handler: router,
	}

	gqlServer := gql.InitGQL(resolverContainer)

	router.Post("/gql", gqlServer.GraphQL())

	g.Add(
		func() error {
			return s.ListenAndServe()
		},
		func(e error) {
			if nil != e {
				panic(e)
			}
		},
	)
}
