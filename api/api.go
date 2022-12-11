package api

import (
	"LIFO/api/todos"
	"LIFO/server"
)

var app = server.GetApp()

// Start ...
func Start() {

	todos.Handle()

}
