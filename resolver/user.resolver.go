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

func (r *userResolver) CreatedAt() (*graphql.Time, error) {
	if r.u.CreatedAt.IsZero() {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.CreatedAt.String())
	return &graphql.Time{Time: t}, err
}

func (r *userResolver) UpdatedAt() (*graphql.Time, error) {
	if r.u.UpdatedAt.IsZero() {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.UpdatedAt.String())
	return &graphql.Time{Time: t}, err
}

func (r *userResolver) Name() *string {
	return &r.u.Name
}

func (r *userResolver) Username() string {
	return r.u.Username
}

func (r *userResolver) Email() string {
	return r.u.Email
}
