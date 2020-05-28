package service

import (
	"time"

	"github.com/mrtrom/go-graphql-example-api/model"
)

const (
	defaultListFetchSize = 10
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetOne() (*model.User, error) {
	user := &model.User{
		ID:        "1",
		CreatedAt: time.Now().String(),
		Email:     "example@email.com",
	}

	return user, nil
}
