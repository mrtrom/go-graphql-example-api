package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/mrtrom/go-graphql-example-api/model"
)

type userResolver struct {
	u *model.User
}

func (r *userResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID)
}

func (r *userResolver) Name() *string {
	return &r.u.Name
}

func (r *userResolver) LastName() *string {
	return &r.u.LastName
}
