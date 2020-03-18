package domain

import "context"

//User - model defnition
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//UserRepository -  repository interface for user
type UserRepository interface {
	Save(context.Context, User) error
	Find(context.Context, string) error
}
