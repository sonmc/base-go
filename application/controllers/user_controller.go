package controller

import (
	"encoding/json"
	"log"
	"net/http"

	service "be/application/services"
) 

type UserController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
}

type userController struct{
	userService service.UserService
	logger *log.Logger
}

func NewUserController(logger *log.Logger, userService service.UserService) UserController {
	return &userController{userService: userService, logger: logger}
}

func (user *userController) GetAll(response http.ResponseWriter, request *http.Request) {

	user.logger.Println("GetUsers controller method called ")

	response.Header().Set("Content-Type", "application/json")
	users, err := user.userService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Internal server error"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}
