package resolver

import (
	"context"
	"fmt"

	"github.com/mrtrom/go-graphql-example-api/config"
	"github.com/mrtrom/go-graphql-example-api/service"
	"go.uber.org/zap"
)

// User resolver
func (r *RootResolver) User(ctx context.Context) (*UserResolver, error) {
	log := ctx.Value(config.CTXLog).(*zap.SugaredLogger)
	user, err := ctx.Value(config.CTXUserService).(*service.UserService).GetOne()
	if err != nil {
		log.Error(fmt.Sprintf("There was an error %s", err))
	}

	return &UserResolver{user}, nil
}
