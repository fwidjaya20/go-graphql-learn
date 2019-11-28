package http

import (
	"encoding/json"
	"github.com/fwidjaya20/go-graphql-learn/internal/gql"
	"github.com/graphql-go/graphql"
	"github.com/payfazz/go-apt/pkg/fazzcommon/response"
	"net/http"
)

type Server struct {
	GqlSchema *graphql.Schema
}

type reqBody struct {
	Query string `json:"query"`
}

func (s *Server) GraphQL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Must provide graphql query in request body", 400)
			return
		}

		var rBody reqBody
		if err := json.NewDecoder(r.Body).Decode(&rBody); err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		result := gql.ExecureQuery(rBody.Query, *s.GqlSchema)

		response.Json(w, result, 200)
	}
}