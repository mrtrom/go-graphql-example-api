package resolver

import (
	"context"
	"fmt"

	"github.com/mrtrom/go-graphql-example-api/config"
	"github.com/mrtrom/go-graphql-example-api/model"
	"github.com/mrtrom/go-graphql-example-api/service"
	"go.uber.org/zap"
)

type newUserArgs struct {
	Email    string
	Username string
}

func (r *RootResolver) CreateUser(ctx context.Context, args newUserArgs) (*UserResolver, error) {
	log := ctx.Value(config.CTXLog).(*zap.SugaredLogger)
	user := &model.User{
		Email:    args.Email,
		Username: args.Username,
	}

	user, err := ctx.Value(config.CTXUserService).(*service.UserService).CreateUser(user)
	if err != nil {
		log.Errorf(fmt.Sprintf("There was an error creating the chat: %s", err))
	}

	return &UserResolver{user}, nil
}
