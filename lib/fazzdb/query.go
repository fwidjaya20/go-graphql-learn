package fazzdb

import (
	"context"
	"github.com/fwidjaya20/go-graphql-learn/config"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

func GetQuery(ctx context.Context) *fazzdb.Query {
	t, err := fazzdb.GetTransactionOrQueryContext(ctx)

	if nil != err {
		dbConn := config.GetDB()
		q := fazzdb.QueryDb(dbConn, config.GetIfQueryConfig(config.I_QUERY_CONFIG))

		return q
	}

	return t
}