package routes

import (
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
	logger     = log.New(os.Stdout, "service-api : ", log.LstdFlags)
	httpRouter = NewMuxRouter(logger)
)

// Run will start the server
func Run() {
	const PORT string = "5000"
	authRoutes()
	userRoutes()
	httpRouter.SERVE(PORT)
}
