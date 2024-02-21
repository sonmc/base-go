package routes

import (
	userCrl "be/application/controllers/user"
)

func userRoutes() {
	httpRouter.GET("/users", userCrl.GetAll)
}
