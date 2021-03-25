package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func(nextHandler http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			authenticated := false

			if authenticated {
				nextHandler(writer, request)
			} else {
				fmt.Println("YOU AREN'T ALLOWED TO ACCESS THE ", request.URL.Path, " RESOURCE!!!")
			}
		}
	}
}
