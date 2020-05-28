package resolver

import (
	"context"
	"fmt"

	"github.com/mrtrom/go-graphql-example-api/service"
)

// Users func
func (r *RootResolver) Users(ctx context.Context) (*userResolver, error) {
	user, err := ctx.Value("userService").(*service.UserService).GetOne()
	if err != nil {
		fmt.Printf("there was an error here %s", err)
	}

	return &userResolver{user}, nil
}
