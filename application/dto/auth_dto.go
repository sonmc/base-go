package dto

import (
	"errors"
)

var (
	ErrUserNameCannotBeBlank       = errors.New("title can not be blank")
	ErrPasswordCannotBeBlank       = errors.New("item not found") 
	ErrAccountIsNotMatch           = errors.New("Invalid username or password")
)
 
type Auth struct {       
	Email    string      
	Password    string      
}
  
func (auth Auth) Validate() error {
	if auth.Email == "" {
		return ErrUserNameCannotBeBlank
	}
	if auth.Password == "" {
		return ErrPasswordCannotBeBlank
	}
	return nil;
}