package main

import "SimpleWebServer/server"

func main() {
	server := server.New()
	server.Server(8080)
}
