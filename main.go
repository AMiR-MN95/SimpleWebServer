package main

import "SimpleWebServer/server"

func main() {
	server := server.Start()
	server.Serve(8080)
}
