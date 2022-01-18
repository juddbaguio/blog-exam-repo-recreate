package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/juddbaguio/blog-exam-repo-recreate/pkg/logic"
)

type Server struct {
	Router *mux.Router
	logic  logic.Logic
}

func InitServer(logic logic.Logic) (*Server, error) {
	router := mux.NewRouter()
	s := &Server{
		Router: router,
		logic:  logic,
	}

	s.InitRoutes()
	return s, nil
}

func (s *Server) InitRoutes() {
	s.SetupArticleRoutes()
}

func (s *Server) Start() error {
	srv := http.Server{
		Addr:         ":3000",
		Handler:      s.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-shutdown
		log.Println("server will start shutting down")

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Println("server shutdown")
		}
	}()

	return srv.ListenAndServe()
}
