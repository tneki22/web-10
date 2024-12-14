package usecase

import "web-10/pkg/vars"

type Provider interface {
	FetchCount() (int, error)
	IncreaseCount(int) error
}

type Usecase struct {
	provider Provider
}

func NewUsecase(provider Provider) *Usecase {
	return &Usecase{provider: provider}
}

func (u *Usecase) FetchCount() (int, error) {
	return u.provider.FetchCount()
}

func (u *Usecase) IncreaseCount(value int) error {
	if value <= 0 {
		return vars.ErrInvalidValue
	}
	return u.provider.IncreaseCount(value)
}
