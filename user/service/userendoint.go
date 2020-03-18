package service

import (
	"context"

	"example.com/go-microservice/user/domain"
	"example.com/go-microservice/user/requests"
	"example.com/go-microservice/user/responses"
	"github.com/go-kit/kit/endpoint"
)

//UserEndpoint - Struct UserEndpoint
type UserEndpoint struct {
	Save endpoint.Endpoint
	Find endpoint.Endpoint
}

//MakeEndpoints - MakeEndpoints for User service
func MakeEndpoints(us UserService) UserEndpoint {
	return UserEndpoint{
		Save: makeSaveUserEndpoint(us),
		Find: makeFindUserEndpoint(us),
	}
}

//makeSaveUserEndpoint - Transform request and call service
func makeSaveUserEndpoint(us UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.UserSaveRequest)
		user := &domain.User{}
		user.FirstName = req.FirstName
		user.LastName = req.LastName
		user.Email = req.Email
		id, err := us.Save(ctx, user)
		return responses.UserSaveResponse{ID: id}, err
	}
}

//makeFindUserEndpoint - Transform request and call service
func makeFindUserEndpoint(us UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(requests.UserFindRequest)
		user, err := us.Find(ctx, req.ID)
		res := &responses.UserFindResponse{}
		res.FirstName = user.FirstName
		res.LastName = user.LastName
		res.Email = user.LastName
		return res, err
	}
}
