// main.go
package main

import (
	"albus-auth/database"
	"albus-auth/routes"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Connect to the database
	_, err := database.ConnectDB()
	if err != nil {
		// If unable to connect, panic
		panic("could not connect to db")
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env: %v", err)
	}

	// Print a success message if connection is successful
	fmt.Println("Connection is successful")

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

	// Start the server
	err = app.Listen(os.Getenv("SERVER_PORT"))
	if err != nil {
		// If unable to start the server, panic
		panic("could not start server")
	}
}
