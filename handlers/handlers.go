package handlers

import (
	"crypto-exchange/constants"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CurrencyPairList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success":    true,
		"currencies": constants.AvailableCurrencyPairs,
	})
}

func CurrencyPairDetail(c *fiber.Ctx) error {
	pair := c.Params("pair")

	return c.JSON(fiber.Map{
		"success":      true,
		"is_available": true,
		"pair":         pair,
		"rate":         0.3337,
	})
}

func CurrencyRate(c *fiber.Ctx) error {
	from := c.FormValue("from")
	to := c.FormValue("to")
	value, err := strconv.ParseFloat(c.FormValue("value"), 64)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "Invalid value.",
		})
	}

	return c.JSON(fiber.Map{
		"success":         true,
		"is_available":    true,
		"from":            from,
		"to":              to,
		"value":           value,
		"converted_value": value * 0.3337,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
