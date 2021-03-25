package main

import (
	"fmt"
	"log"
	"net/http"
)

func CheckAuth() Middleware {
	return func(nextHandler http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			authenticated := true

			if authenticated {
				nextHandler(writer, request)
			} else {
				fmt.Println("YOU AREN'T ALLOWED TO ACCESS THE ", request.URL.Path, " RESOURCE!!!")
			}
		}
	}
}

func LogStuff() Middleware {
	return func(nextHandler http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			log.Println("Going to serve the request on ", request.URL.Path)

			nextHandler(writer, request)
		}
	}
}
