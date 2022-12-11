package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"net"
	"os"
)

func GetAddress() string {
	return fmt.Sprintf("%s:%s", GetHost(), GetPort())
}

func newServer() *Server {
	listener, err := net.Listen("tcp4", GetAddress())
	if err != nil {
		fmt.Println(err)
	}

	return &Server{listener}
}

func newFiberApp() *fiber.App {
	newApp := fiber.New()
	newApp.Use(cors.New())
	newApp.Use(recover.New())
	return newApp
}

func GetApp() *fiber.App {
	return app
}

func GetHost() string {
	host := os.Getenv("HOST")

	if host == "" {
		return "localhost"
	}

	return host
}

func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return "3000"
	}

	return port
}
