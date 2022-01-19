package pkg

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oguzhantasimaz/arog-gora-quote-api/internal"
)

func Get(c *fiber.Ctx) error {
	quote := *internal.GetRandomQuote()
	fmt.Println(quote.Text)
	return c.Status(200).JSON(quote)
}
