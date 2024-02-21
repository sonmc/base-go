package userFlow

import (
	"errors"
)

var ErrFindOneFailed = errors.New("404 not found")
var ErrAuthenticationFailed = errors.New("authentication failed")
