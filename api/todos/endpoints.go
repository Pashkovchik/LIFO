package todos

import (
	"LIFO/server"
	"fmt"
)

var app = server.GetApp().Group("/api/todos")

var count = 10
var todos = make([]*Todo, count)

func Handle() {
	h := handlers{}

	testData()

	app.Get("/", h.Todos)
	app.Post("/", h.Push)
	app.Delete("/", h.Pop)
}

func testData() {

	for i := 0; i < count; i++ {
		id := fmt.Sprintf("%d", currentId)
		todo := &Todo{
			ID:    id,
			Name:  "name",
			Value: "todo from api",
		}

		currentId++

		todos[i] = todo

	}
}
