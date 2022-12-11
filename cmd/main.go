package main

import (
	"LIFO/api"
	"LIFO/server"
)

func main() {
	api.Start()
	server.Start()
}
