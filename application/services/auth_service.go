package service

import (
	"be/application/dto"
	repository "be/application/repositories"

	model "be/domain/user"
)

type AuthService interface {
	Login(authDto dto.Auth) (model.User, error)
}

type authService struct{
	repository repository.UserRepository
}


func NewAuthService(repository repository.UserRepository) AuthService {
	return &authService{repository: repository}
}

func (service *authService) Login(auth dto.Auth) (model.User, error) {
	return  model.User{}, nil;
}
