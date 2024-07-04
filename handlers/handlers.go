package handlers

import (
	"crypto-exchange/database"
	"crypto-exchange/workers"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type CurrencyPair struct {
	Pair        string  `json:"pair_name"`
	IsAvailable bool    `json:"is_available"`
	Rate        float64 `json:"rate"`
}

func CurrencyPairsList(c *fiber.Ctx) error {
	rawPairs := database.GetCurrencyPairs()

	var pairs []CurrencyPair
	for _, pair := range rawPairs {
		pairs = append(pairs, CurrencyPair{
			Pair:        strings.Replace(pair.Name, "/", "-", -1),
			IsAvailable: true,
			Rate:        pair.Rate,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"paris":   pairs,
	})
}

func CurrencyPairDetail(c *fiber.Ctx) error {
	pair := c.Params("pair")

	rawPair := database.GetCurrencyPair(strings.Replace(pair, "-", "/", -1))

	if rawPair.ID == 0 {
		return c.JSON(fiber.Map{
			"success":      false,
			"is_available": false,
			"error":        "Currency pair is not found.",
		})
	}

	return c.JSON(fiber.Map{
		"success":      true,
		"is_available": true,
		"pair":         rawPair.Name,
		"rate":         rawPair.Rate,
	})
}

func CurrencyPairRate(c *fiber.Ctx) error {
	from := c.FormValue("from")
	to := c.FormValue("to")
	value, err := strconv.ParseFloat(c.FormValue("amount"), 64)
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "Invalid amount value.",
		})
	}

	rawPair := database.GetCurrencyPair(from + "/" + to)
	if rawPair.ID == 0 {
		return c.JSON(fiber.Map{
			"success":      false,
			"is_available": false,
			"error":        "Currency pair is not found.",
		})
	}

	workers.SignalCurrencyUpdater(from + "/" + to)

	return c.JSON(fiber.Map{
		"success":      true,
		"is_available": true,
		"from":         from,
		"to":           to,
		"amount":       value * rawPair.Rate,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
