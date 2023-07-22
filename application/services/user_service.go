package service

import (
	repository "be/application/repositories"

	model "be/domain/user"
)

type UserService interface {
	FindAll() ([]model.User, error)
}

type service struct{
	repository repository.UserRepository
}


func NewUserService(repository repository.UserRepository) UserService {
	return &service{repository: repository}
}

func (service *service) FindAll() ([]model.User, error) {
	return service.repository.FindAll()
}
