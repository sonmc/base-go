package routes

import (
	authCtrl "be/application/controllers/auth"
)

func authRoutes() {
	httpRouter.POST("/auth/login", authCtrl.Login)
}
