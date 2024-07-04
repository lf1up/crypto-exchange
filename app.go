package main

import (
	"crypto-exchange/database"
	"crypto-exchange/handlers"
	"crypto-exchange/workers"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
	dev  = flag.Bool("dev", false, "Enable development mode")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.Connect(*dev)

	// Start the background worker
	go workers.StartCurrencyUpdater()
	go workers.ScheduleBackgroundUpdate(1) // 1 minute

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/pairs", handlers.CurrencyPairList)
	v1.Get("/pairs/:pair", handlers.CurrencyPairDetail)
	v1.Post("/convert", handlers.CurrencyPairPrice)

	// Setup static files [disabled because of no need for static frontend files in this app]
	// app.Static("/", "./static/public")

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
