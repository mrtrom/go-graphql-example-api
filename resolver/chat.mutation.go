package resolver

import (
	"context"
	"fmt"

	"github.com/mrtrom/go-graphql-example-api/config"
	"github.com/mrtrom/go-graphql-example-api/model"
	"github.com/mrtrom/go-graphql-example-api/service"
	"go.uber.org/zap"
)

type newChatArgs struct {
	Content string
	From    string
}

func (r *RootResolver) CreateChat(ctx context.Context, args newChatArgs) (*ChatResolver, error) {
	log := ctx.Value(config.CTXLog).(*zap.SugaredLogger)
	chat := &model.Chat{
		Content: args.Content,
		From:    args.From,
	}

	chat, err := ctx.Value(config.CTXChatService).(*service.ChatService).CreateChat(chat)
	if err != nil {
		log.Errorf(fmt.Sprintf("There was an error creating the chat: %s", err))
	}

	return &ChatResolver{chat}, nil
}
