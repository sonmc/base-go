package repositories

import (
	model "be/domain/user"
	database "be/infrastructure"
	"log"
) 
 

type UserRepository interface {
	//Save(product *model.Product) (*model.Product, error)
	FindAll() ([]model.User, error)
}

type repo struct {
	logger *log.Logger
}

func NewUserRepository(logger *log.Logger) UserRepository {
	return &repo{logger: logger}
}

//func (*repo) Save(product *model.Product) (*model.Product, error)  {
//	return product, nil
//}

func (repo *repo) FindAll() ([]model.User, error)  { 
	repo.logger.Println("Fetching all users from database") 
	var products []model.User 
	database.DB.Find(&products) 
	return products, nil
}