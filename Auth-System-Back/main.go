// main.go
package main

import (
	"albus-auth/database"
	"albus-auth/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect to the database
	_, err := database.ConnectDB()
	if err != nil {
		// If unable to connect, panic
		panic("could not connect to db")
	}

	// Print a success message if connection is successful
	fmt.Println("Connection is successful")

	// Initialize Fiber app
	app := fiber.New()

	// Adding CORS middleware with specific origin
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // Replace with your frontend URL
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization,Accept,Origin",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Setup routes
	routes.SetUpRoutes(app)

	// Start the server
	err = app.Listen(":8000")
	if err != nil {
		// If unable to start the server, panic
		panic("could not start server")
	}
}
