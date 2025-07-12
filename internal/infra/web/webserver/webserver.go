package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type (
	WebServer struct {
		Router        chi.Router
		Handlers      []HandlerData
		WebServerPort string
	}

	HandlerData struct {
		Method string
		Path   string
		Handler http.HandlerFunc
	}
)

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make([]HandlerData, 0),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, HandlerData{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		s.Router.Method(handler.Method, handler.Path, handler.Handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
