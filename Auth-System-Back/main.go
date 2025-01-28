// main.go
package main

import (
	"albus-auth/database"
	"albus-auth/routes"
	"albus-auth/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Load environment variables
	config, err := utils.LoadConfig()

	// Initialize logger with the log level from environment variables
	err = utils.Initialize(config.LogLevel)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer utils.Sync()

	sugar := utils.SugaredLogger

	sugar.Infow("Starting application with config",
		"DBHost", config.DBHost,
		"DBPort", config.DBPort,
		"FrontendURL", config.FrontendURL,
		"ServerPort", config.ServerPort,
	)

	// Connect to the database
	_, err = database.ConnectDB(config, sugar)
	if err != nil {
		sugar.Fatalw("Failed to connect to database",
			"error", err,
		)
	}
	sugar.Info("Successfully connected to database")

	// Initialize Fiber app
	app := fiber.New()

	// Adding CORS middleware with specific origin
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("FRONTEND_URL"),
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization,Accept,Origin",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Setup routes
	routes.SetUpRoutes(app)

	sugar.Infof("Server starting on port%s", config.ServerPort)
	if err := app.Listen(config.ServerPort); err != nil {
		sugar.Fatalw("Server failed to start",
			"error", err,
		)
	}
}
