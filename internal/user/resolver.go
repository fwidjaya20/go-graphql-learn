package user

import (
	"context"
	"errors"
	"github.com/fwidjaya20/go-graphql-learn/internal/user/repositories/queries"
	"github.com/graphql-go/graphql"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

type IResolver interface {
	Resolve(param graphql.ResolveParams) (interface{}, error)
}

type resolver struct {
	query queries.UserQueryInterface
}

func NewUserResolver() IResolver {
	return &resolver{
		query: queries.NewUserQuery(),
	}
}

func (ur *resolver) Resolve(param graphql.ResolveParams) (interface{}, error) {
	fullName, ok := param.Args["full_name"].(string)

	if !ok {
		return nil, errors.New("params 'Full Name' not found")
	}

	users, err := ur.query.Find(context.Background(), []fazzdb.SliceCondition{
		{
			Connector:  fazzdb.CO_NONE,
			Field:      "full_name",
			Operator:   fazzdb.OP_ILIKE,
			Value:      "%" + fullName + "%",
			Conditions: nil,
		},
	})

	if nil != err {
		return nil, errors.New(err.Error())
	}

	if len(users) < 1 {
		return nil, errors.New("no user found")
	}

	return users, nil
}
