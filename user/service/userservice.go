package service

import (
	"context"

	"example.com/go-microservice/user/domain"
)

//UserService the interface
type UserService interface {
	Save(context.Context, *domain.User) (string, error)
	Find(context.Context, string) (*domain.User, error)
}

//UserServiceImpl Implementation struct
type UserServiceImpl struct {
}

//Save - Method to save user
func (us *UserServiceImpl) Save(ctx context.Context, user *domain.User) (string, error) {
	return "ID123", nil
}

//Find - Method to find a user based on id
func (us *UserServiceImpl) Find(ctx context.Context, id string) (*domain.User, error) {
	return nil, nil
}
