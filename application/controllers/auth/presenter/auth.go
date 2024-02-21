package presenter

import (
	"errors"
)

var (
	ErrUserNameCannotBeBlank = errors.New("title can not be blank")
	ErrPasswordCannotBeBlank = errors.New("item not found")
)

type Auth struct {
	Username string
	Password string
}

func (auth Auth) Validate() error {
	if auth.Username == "" {
		return ErrUserNameCannotBeBlank
	}
	if auth.Password == "" {
		return ErrPasswordCannotBeBlank
	}
	return nil
}
