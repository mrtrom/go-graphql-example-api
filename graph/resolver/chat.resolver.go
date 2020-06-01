package resolver

import (
	"context"
	"fmt"

	"github.com/mrtrom/go-graphql-example-api/config"
	"github.com/mrtrom/go-graphql-example-api/graph/model"
	"github.com/mrtrom/go-graphql-example-api/service"
	"go.uber.org/zap"
)

type chatResolver struct{ *Resolver }

func (r *queryResolver) Chats(ctx context.Context) ([]*model.Chat, error) {
	log := ctx.Value(config.CTXLog).(*zap.SugaredLogger)

	chats, err := ctx.Value(config.CTXChatService).(*service.ChatService).GetAll()
	if err != nil {
		log.Error(fmt.Sprintf("There was an error %s", err))
	}

	return chats, nil
}

func (r *mutationResolver) CreateChat(ctx context.Context, content string, from string) (*model.Chat, error) {
	log := ctx.Value(config.CTXLog).(*zap.SugaredLogger)
	chat := &model.Chat{
		Content: content,
		From:    from,
	}

	chat, err := ctx.Value(config.CTXChatService).(*service.ChatService).CreateChat(chat)
	if err != nil {
		log.Errorf(fmt.Sprintf("There was an error creating the chat: %s", err))
	}

	return chat, nil
}
