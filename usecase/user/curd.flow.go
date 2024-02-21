package userFlow

import "be/application/controllers/auth/presenter"

func List() (presenter.Auth, error) {
	return presenter.Auth{}, ErrFindOneFailed
}
