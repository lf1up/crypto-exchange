package handlers

import (
	"crypto-exchange/constants"

	"github.com/gofiber/fiber/v2"
)

func CurrencyList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success":    true,
		"currencies": constants.AvailableCurrencyPairs,
	})
}

func CurrencyDetail(c *fiber.Ctx) error {
	pair := c.Params("pair")

	return c.JSON(fiber.Map{
		"success":      true,
		"is_available": true,
		"pair":         pair,
		"rate":         0.3337,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
