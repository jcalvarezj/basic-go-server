# Basic Golang Web server

A simple example of a Web server in Golang using the default net/http server and router

Practice of Middeware implementation and JSON serving with Encode/Decode

The /api/users resource serves GET and POST requests. The latter expects a JSON with the fields **id** (int), **name** (string), and **age** (int)

Run with `go build && ./basic-go-server`
