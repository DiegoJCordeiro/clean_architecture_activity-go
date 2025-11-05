package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// WebServer representa o servidor web
type WebServer struct {
	Router *chi.Mux
	Port   string
}

// NewWebServer cria um novo servidor web
func NewWebServer(port string) *WebServer {
	return &WebServer{
		Router: chi.NewRouter(),
		Port:   port,
	}
}

// AddMiddleware adiciona middlewares ao servidor
func (s *WebServer) AddMiddleware() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
}

// Start inicia o servidor
func (s *WebServer) Start() error {
	fmt.Printf("🚀 REST API rodando na porta %s\n", s.Port)
	return http.ListenAndServe(":"+s.Port, s.Router)
}
