package main

import (
	"log"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		log.Fatal("unable to run server", err)
		return err
	}
	log.Print("server running on", s.port)
	return nil
}

func (s *Server) Handle(path string, method string, handler http.HandlerFunc) {
	//Check if the path already exists
	if !s.router.FindPath(path) {
		//If not path then create a new one
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

type Handler func(w http.ResponseWriter, r *http.Request)

type Middleware func(http.HandlerFunc) http.HandlerFunc
