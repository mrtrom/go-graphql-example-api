package service

import (
	"github.com/jinzhu/gorm"
	"github.com/mrtrom/go-graphql-example-api/model"
	"go.uber.org/zap"
)

type ChatService struct {
	db  *gorm.DB
	log *zap.SugaredLogger
}

func NewChatService(db *gorm.DB, log *zap.SugaredLogger) *ChatService {
	return &ChatService{db: db, log: log}
}

func (c *ChatService) GetAll() ([]*model.Chat, error) {
	chats := make([]*model.Chat, 0)

	c.db.Find(&chats)

	return chats, nil
}

func (c *ChatService) CreateChat(chat *model.Chat) (*model.Chat, error) {
	if err := c.db.Create(&chat).Error; err != nil {
		return nil, err
	}

	return chat, nil
}
