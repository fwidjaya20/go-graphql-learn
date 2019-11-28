package queries

import (
	"context"
	"github.com/fwidjaya20/go-graphql-learn/internal/user/models"
	fazzdbLib "github.com/fwidjaya20/go-graphql-learn/lib/fazzdb"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

type UserQueryInterface interface {
	Find(context context.Context, conditions []fazzdb.SliceCondition) ([]*models.User, error)
}

type query struct {
	User *models.User
}

func NewUserQuery() UserQueryInterface {
	return &query{
		User: models.UserModel(),
	}
}

func (qr *query) Find(ctx context.Context, conditions []fazzdb.SliceCondition) ([]*models.User, error) {
	q := fazzdbLib.GetQuery(ctx)

	users, err := q.Use(qr.User).WhereMany(conditions...).AllCtx(ctx)

	return users.([]*models.User), err
}
