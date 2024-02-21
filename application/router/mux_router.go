package routes

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type MuxRouter struct {
	logger *log.Logger
}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter(logger *log.Logger) Router {
	return &MuxRouter{logger: logger}
}

func (*MuxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*MuxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (muxRouter *MuxRouter) SERVE(port string) {

	muxRouter.logger.Println("HTTP server started on port", port)

	// Setting up the web server
	s := &http.Server{
		Addr:              ":" + port,
		Handler:           muxDispatcher,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	// GoRoutines
	go func() {
		// Running the web server
		err := s.ListenAndServe()
		if err != nil {
			muxRouter.logger.Println("Server shutting down on port ", port)
			os.Exit(1)
		}
	}()

	// Creating a channel
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig := <-c
	muxRouter.logger.Println("Receive signal:", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shut down the server context
	_ = s.Shutdown(ctx)
}
