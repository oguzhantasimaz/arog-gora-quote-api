package main

import (
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/oguzhantasimaz/arog-gora-quote-api/pkg/handlers"
)

func main() {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = "0.0.0.0"
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "5000"
	}

	listenAddress := net.JoinHostPort(host, port)

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", handlers.Get)

	app.Listen(listenAddress)
}
