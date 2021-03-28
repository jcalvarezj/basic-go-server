package main

import (
	"fmt"
)

func main() {
	fmt.Println("Server running on port 8080")
	server := NewServer(":8080")

	APIHandler := server.AddMiddlewares(HandleAPIHome, CheckAuth(), LogStuff())

	server.Handle("/", HandleRoot)
	server.Handle("/api", APIHandler)
	server.Handle("/api/users", HandleUsersPost)
	server.Listen()
}
