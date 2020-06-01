package resolver

import (
	"context"
	"fmt"

	"github.com/mrtrom/go-graphql-example-api/config"
	"github.com/mrtrom/go-graphql-example-api/graph/model"
	"github.com/mrtrom/go-graphql-example-api/service"
	"go.uber.org/zap"
)

type userResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	log := ctx.Value(config.CTXLog).(*zap.SugaredLogger)
	user, err := ctx.Value(config.CTXUserService).(*service.UserService).GetOne()
	if err != nil {
		log.Error(fmt.Sprintf("There was an error %s", err))
	}

	return user, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, email string, username string) (*model.User, error) {
	log := ctx.Value(config.CTXLog).(*zap.SugaredLogger)
	user := &model.User{
		Email:    email,
		Username: username,
	}

	user, err := ctx.Value(config.CTXUserService).(*service.UserService).CreateUser(user)
	if err != nil {
		log.Errorf(fmt.Sprintf("There was an error creating the chat: %s", err))
	}

	return user, nil
}
