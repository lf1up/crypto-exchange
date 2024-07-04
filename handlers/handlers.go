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

// CurrencyPairsList retrieves a list of all available currency pairs from the database.
// It returns a JSON response containing an array of currency pairs.
// Each currency pair includes its name, conversion rate, and other relevant details.
//
// @Summary List of currency pairs
// @Description Retrieves a list of all available currency pairs.
// @Tags currency
// @Accept json
// @Produce json
// @Success 200 {array} map[string]interface{} "Success response with list of currency pairs"
// @Router /currencies [get]
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

// CurrencyPairDetail retrieves the detailed information of a specific currency pair.
// It extracts the 'pair' parameter from the URL path, searches for the currency pair details in the database,
// and returns a JSON response with the detailed information.
// If the currency pair is not found, it returns a JSON response indicating failure.
//
// @Summary Currency pair details
// @Description Retrieves detailed information of a specific currency pair.
// @Tags currency
// @Accept json
// @Produce json
// @Param pair path string true "Currency pair code"
// @Success 200 {object} map[string]interface{} "Success response with currency pair details"
// @Failure 404 {object} map[string]interface{} "Error response when currency pair is not found"
// @Router /currencies/{pair} [get]
func CurrencyPairDetail(c *fiber.Ctx) error {
	pair := c.Params("pairName")

	pairName := strings.Replace(pair, "-", "/", -1)
	rawPair := database.GetCurrencyPair(pairName)

	if rawPair.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
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

// GetCurrencyPair handles the request for converting currency from one type to another.
// It first constructs a pair name from the 'from' and 'to' query parameters,
// then retrieves the currency pair information from the database.
// If the currency pair is not found, it returns a JSON response indicating failure.
// Otherwise, it signals the currency updater and returns a JSON response with the conversion result.
//
// @Summary Convert currency
// @Description Converts currency from one type to another using the specified amount.
// @Tags currency
// @Accept json
// @Produce json
// @Param from query string true "Currency code to convert from"
// @Param to query string true "Currency code to convert to"
// @Param amount query float64 true "Amount to convert"
// @Success 200 {object} map[string]interface{} "Success response with conversion result"
// @Failure 400 {object} map[string]interface{} "Error response when currency pair is not found or amount is invalid"
// @Router /currencies/convert [post]
func CurrencyPairRate(c *fiber.Ctx) error {
	from := c.FormValue("from")
	to := c.FormValue("to")
	value, err := strconv.ParseFloat(c.FormValue("amount"), 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid amount value.",
		})
	}

	pairName := from + "/" + to
	rawPair := database.GetCurrencyPair(pairName)

	if rawPair.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
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

// APIStatus returns the current status of the API.
// It simply returns a JSON response indicating that the API is operational.
//
// @Summary API status
// @Description Checks the current status of the API.
// @Tags status
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Success response indicating API is OK"
// @Router / [get]
func APIStatus(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"status":  "OK",
	})
}

// NotFound handles requests to undefined routes.
// It returns a custom 404 error page to the client.
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
