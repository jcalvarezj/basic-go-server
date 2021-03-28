package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

// Serves /
func HandleRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "<h1>Hello World!</h1>")
}

// Serves /api
func HandleAPIHome(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "<h1>API HOME!!</h1>")
}

// Serves /api/users. GET to show a JSON, and POST to store and return it
func HandleUsersPost(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
		case "GET":
			user := User{ 123456, "Pepe", 23 }
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(&user)
		case "POST":
			decoder := json.NewDecoder(request.Body)
			var user User
			err := decoder.Decode(&user)
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
			} else {
				fmt.Fprintln(writer, "The user is ", user)
			}
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
