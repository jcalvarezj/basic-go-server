package main

import (
	"net/http"
)

type Server struct {
	port string
	router *Router
}

func NewServer(port string) *Server{
	return &Server{port, NewRouter()}
}

func (self *Server) Handle(path string, handler http.HandlerFunc) {
	self.router.rules[path] = handler
}

func (self *Server) Listen() error {
	http.Handle("/", self.router)
	return http.ListenAndServe(self.port, nil)
}

func (self *Server) AddMiddlewares(baseHandler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	encapsulatedHandler := baseHandler

	for _, middleware := range middlewares {
		encapsulatedHandler = middleware(encapsulatedHandler)
	}

	return encapsulatedHandler
}
