package service

import (
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
		ID:       1,
		Name:     "Name",
		LastName: "Last name",
	}

	return user, nil
}
