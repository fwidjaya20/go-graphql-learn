package containers

import "github.com/fwidjaya20/go-graphql-learn/internal/user"

type ResolverContainer struct {
	UserResolver user.IResolver
}

func NewResolverContainer() ResolverContainer {
	return ResolverContainer{
		UserResolver: user.NewUserResolver(),
	}
}
