package service

import (
	"github.com/jinzhu/gorm"
	"github.com/mrtrom/go-graphql-example-api/model"
	"go.uber.org/zap"
)

type UserService struct {
	db  *gorm.DB
	log *zap.SugaredLogger
}

func NewUserService(db *gorm.DB, log *zap.SugaredLogger) *UserService {
	return &UserService{db: db, log: log}
}

func (u *UserService) GetOne() (*model.User, error) {
	user := &model.User{}

	u.db.First(&user, 1)

	return user, nil
}
