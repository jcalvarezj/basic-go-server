package main

import (
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router {make(map[string]http.HandlerFunc)}
}

func (self *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exists := self.rules[path]
	return handler, exists
}

func (self *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler, exists := self.FindHandler(request.URL.Path)

	if exists {
		handler(writer, request)
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}
