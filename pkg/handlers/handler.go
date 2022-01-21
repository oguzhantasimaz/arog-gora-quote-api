package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oguzhantasimaz/arog-gora-quote-api/internal/models"
)

func Get(c *fiber.Ctx) error {
	quote := *models.GetRandomQuote()
	return c.Status(200).JSON(quote)
}
