package vars

import "errors"

var (
	ErrAlreadyExist = errors.New("already exist")
	ErrInvalidValue = errors.New("value must be greater than 0")
)
