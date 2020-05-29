package resolver

import (
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/mrtrom/go-graphql-example-api/model"
)

type UserResolver struct {
	u *model.User
}

func (r *UserResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID)
}

func (r *UserResolver) CreatedAt() (*graphql.Time, error) {
	if r.u.CreatedAt.IsZero() {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.CreatedAt.String())
	return &graphql.Time{Time: t}, err
}

func (r *UserResolver) UpdatedAt() (*graphql.Time, error) {
	if r.u.UpdatedAt.IsZero() {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.UpdatedAt.String())
	return &graphql.Time{Time: t}, err
}

func (r *UserResolver) Name() *string {
	return &r.u.Name
}

func (r *UserResolver) Username() string {
	return r.u.Username
}

func (r *UserResolver) Email() string {
	return r.u.Email
}
