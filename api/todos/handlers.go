package todos

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

var currentId int

type handlers struct{}

type AddToDoRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Todo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (h *handlers) Todos(c *fiber.Ctx) error {
	time.Sleep(time.Millisecond * 1000)
	return c.JSON(todos)
}

func (h *handlers) Push(c *fiber.Ctx) error {
	u := new(AddToDoRequest)

	// Parse body into struct
	if err := c.BodyParser(&u); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	id := fmt.Sprintf("%d", currentId)
	newToDo := &Todo{
		ID:    id,
		Name:  u.Name,
		Value: u.Value,
	}
	todos = append(todos, newToDo)

	currentId++

	return c.SendString("Add is good")
}

func (h *handlers) Pop(c *fiber.Ctx) error {
	if currentId == 0 {
		return c.Status(400).SendString("Stack is empty")
	}
	todos = todos[:currentId-1]
	currentId--

	return c.SendString("Delete is completed")
}
