package service

import (
	"github.com/jinzhu/gorm"
	"github.com/mrtrom/go-graphql-example-api/graph/model"
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

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
