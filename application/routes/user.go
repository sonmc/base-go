package routes

import (
	"net/http"
)

func userRoutes() { 
	httpRouter.GET("/users", func(response http.ResponseWriter, request *http.Request) {
		logger.Println(response, "Up and running...")
	}) 
	httpRouter.GET("", userController.GetAll)
	
}
