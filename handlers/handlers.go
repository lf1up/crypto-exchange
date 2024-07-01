package handlers

import (
	"crypto-exchange/constants"
	"crypto-exchange/database"
	"crypto-exchange/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// UserList returns a list of users
func UserList(c *fiber.Ctx) error {
	users := database.Get()

	return c.JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}

// UserCreate registers a user
func UserCreate(c *fiber.Ctx) error {
	user := &models.User{
		// Note: when writing to external database,
		// we can simply use - Name: c.FormValue("user")
		Name: utils.CopyString(c.FormValue("user")),
	}
	database.Insert(user)

	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

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
