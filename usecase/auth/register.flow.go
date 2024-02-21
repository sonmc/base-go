package authFlow

import (
	"be/application/controllers/auth/presenter"
)

func Register(username string, password string) (presenter.Auth, error) {
	if username == "admin" && password == "password" {
		result := presenter.Auth{}
		return result, nil
	}
	return presenter.Auth{}, ErrAuthenticationFailed
}
