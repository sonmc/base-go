package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"be/application/dto"
	service "be/application/services"
) 

type AuthController interface {
	Login(response http.ResponseWriter, request *http.Request)
}

type authController struct{
	authService service.AuthService
	logger *log.Logger
}

func NewAuthController(logger *log.Logger, authService service.AuthService) AuthController {
	return &authController{authService: authService, logger: logger}
}

func (auth *authController) Login(response http.ResponseWriter, request *http.Request) {

	auth.logger.Println("Login controller method called ") 
	response.Header().Set("Content-Type", "application/json")
	var authDto dto.Auth
	json.NewDecoder(request.Body).Decode(&authDto)
	user, err := auth.authService.Login(authDto)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Internal server error"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}
