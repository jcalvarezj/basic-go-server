package main

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}
