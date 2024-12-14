package usecase

import "web-10/internal/query/model"

type Provider interface {
	GetUser(name string) (*model.User, error)
	AddUser(name string) error
}

type Usecase struct {
	provider Provider
}

func NewUsecase(provider Provider) *Usecase {
	return &Usecase{provider: provider}
}

func (u *Usecase) GetUser(name string) (*model.User, error) {
	return u.provider.GetUser(name)
}

func (u *Usecase) AddUser(name string) error {
	return u.provider.AddUser(name)
}
