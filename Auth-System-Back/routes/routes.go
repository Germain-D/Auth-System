// routes/routes.go

package routes

import (
	"albus-auth/controllers" // Replace "your-module-name" with the actual module name

	"github.com/gofiber/fiber/v2"
)

// SetUpRoutes sets up all the routes for the application
func SetUpRoutes(app *fiber.App) {

	// Test route to verify application setup
	app.Get("/", controllers.Hello)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/simplelogin", controllers.SimpleLogin)
	app.Post("/logout", controllers.Logout)

	// google oauth
	app.Get("/auth/google/callback", controllers.GoogleCallback)

	// linkedin oauth
	app.Get("/auth/linkedin/callback", controllers.LinkedInCallback)

	// github oauth
	app.Get("/auth/github/callback", controllers.GitHubCallback)

	// facebook oauth
	app.Get("/auth/facebook/callback", controllers.FacebookCallback)
}
