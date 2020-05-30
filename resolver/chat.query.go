package resolver

import (
	"context"
	"fmt"

	"github.com/mrtrom/go-graphql-example-api/config"
	"github.com/mrtrom/go-graphql-example-api/service"
	"go.uber.org/zap"
)

// Chats method
func (r *RootResolver) Chats(ctx context.Context) (*[]*ChatResolver, error) {
	log := ctx.Value(config.CTXLog).(*zap.SugaredLogger)

	chats, err := ctx.Value(config.CTXChatService).(*service.ChatService).GetAll()
	if err != nil {
		log.Error(fmt.Sprintf("There was an error %s", err))
	}

	var resolvers = make([]*ChatResolver, 0)

	for _, chat := range chats {
		resolvers = append(resolvers, &ChatResolver{chat})
	}

	return &resolvers, nil
}
