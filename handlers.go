package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "<h1>Hello World!</h1>")
}
