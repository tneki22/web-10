package api

import "web-10/internal/query/model"

type Usecase interface {
	GetUser(name string) (*model.User, error)
	AddUser(name string) error
}
