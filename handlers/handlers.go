package handlers

import (
	"crypto-exchange/database"
	"crypto-exchange/workers"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type CurrencyPair struct {
	Name        string  `json:"name"`
	IsAvailable bool    `json:"is_available"`
	Rate        float64 `json:"rate"`
}

func CurrencyPairsList(c *fiber.Ctx) error {
	rawPairs := database.GetCurrencyPairs()

	var pairs []CurrencyPair
	for _, pair := range rawPairs {
		pairs = append(pairs, CurrencyPair{
			Name:        strings.Replace(pair.Name, "/", "-", -1),
			IsAvailable: true,
			Rate:        pair.Rate,
		})
	}

	return c.JSON(fiber.Map{
		"success":        true,
		"currency_pairs": pairs,
	})
}

func CurrencyPairDetail(c *fiber.Ctx) error {
	pair := c.Params("pairName")

	pairName := strings.Replace(pair, "-", "/", -1)
	rawPair := database.GetCurrencyPair(pairName)

	if rawPair.ID == 0 {
		return c.JSON(fiber.Map{
			"success":      false,
			"is_available": false,
			"error":        "Currency pair is not found.",
		})
	}

	workers.SignalCurrencyUpdater(pairName)

	return c.JSON(fiber.Map{
		"success":      true,
		"is_available": true,
		"pair_name":    rawPair.Name,
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

	pairName := from + "/" + to
	rawPair := database.GetCurrencyPair(pairName)

	if rawPair.ID == 0 {
		return c.JSON(fiber.Map{
			"success":      false,
			"is_available": false,
			"error":        "Currency pair is not found.",
		})
	}

	workers.SignalCurrencyUpdater(pairName)

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
