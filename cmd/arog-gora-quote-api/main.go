package main

import (
	"net"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/oguzhantasimaz/arog-gora-quote-api/pkg"
)

func main() {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = "localhost"
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "5000"
	}

	listenAddress := net.JoinHostPort(host, port)

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", pkg.Get)

	app.Listen(listenAddress)
}
