package routes

import (
	controller "be/application/controllers"
	repository "be/application/repositories"
	service "be/application/services"
	"log"
	"net/http"
	"os"
)
 
type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}
var ( 
	logger 			  = log.New(os.Stdout, "user-service-api : ", log.LstdFlags) 
	userRepository    = repository.NewUserRepository(logger)
	userService       = service.NewUserService(userRepository)
 	userController    = controller.NewUserController(logger, userService)
	httpRouter        = NewMuxRouter(logger)
)
// Run will start the server
func Run() {
	const PORT string = "8080"
	authRoutes() 
	userRoutes() 
	httpRouter.SERVE(PORT)
}
