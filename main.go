package main

import (
	"fmt"
)

func main() {
	fmt.Println("Server running on port 8080")
	server := NewServer(":8080")
	server.Handle("/", HandleRoot)
	server.Listen()
}
