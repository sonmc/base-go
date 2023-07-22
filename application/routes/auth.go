package routes

import (
	controller "be/application/controllers"
	service "be/application/services"
)

var (  
	authService       = service.NewAuthService(userRepository)
 	authController    = controller.NewAuthController(logger, authService) 
)

func authRoutes() { 
	httpRouter.POST("/auth/login", authController.Login)
}