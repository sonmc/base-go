package userFlow

import "be/application/controllers/auth/presenter"

func ChangePassword(username string, password string) (presenter.Auth, error) {
	if username == "admin" && password == "password" {
		result := presenter.Auth{
			// Set giá trị cho các trường của struct Auth
		}
		return result, nil
	}
	return presenter.Auth{}, ErrFindOneFailed
}