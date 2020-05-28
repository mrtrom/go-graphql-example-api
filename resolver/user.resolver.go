package resolver

import (
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/mrtrom/go-graphql-example-api/model"
)

type userResolver struct {
	u *model.User
}

func (r *userResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID)
}

func (r *userResolver) Email() *string {
	return &r.u.Email
}

func (r *userResolver) CreatedAt() (*graphql.Time, error) {
	if r.u.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.CreatedAt)
	return &graphql.Time{Time: t}, err
}
