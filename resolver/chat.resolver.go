package resolver

import (
	"strconv"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/mrtrom/go-graphql-example-api/model"
)

type ChatResolver struct {
	c *model.Chat
}

func (r *ChatResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(int(r.c.ID)))
}

func (r *ChatResolver) From() *string {
	return &r.c.From
}

func (r *ChatResolver) Content() *string {
	return &r.c.Content
}

func (r *ChatResolver) CreatedAt() (*graphql.Time, error) {
	if r.c.CreatedAt.IsZero() {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.c.CreatedAt.String())
	return &graphql.Time{Time: t}, err
}
