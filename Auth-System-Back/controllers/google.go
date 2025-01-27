// controllers/google.go
package controllers

import (
	"albus-auth/database"
	"albus-auth/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

/*
func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err) // Handle error appropriately
	}
	return hex.EncodeToString(bytes)
}*/

func GoogleCallback(c *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env: %v", err)
	}

	var (
		googleOauthConfig = &oauth2.Config{
			RedirectURL:  "http://localhost:8000/auth/google/callback",
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		}
		//randomState = generateRandomString(32) // Store the state globally
	)

	// Retrieve the state and code from the query parameters
	state := c.Query("state")
	code := c.Query("code")

	fmt.Println("State:", state)
	fmt.Println("Code:", code)

	// Validate the state
	if state == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing state parameter",
		})
	}

	fmt.Println(googleOauthConfig.ClientID)
	fmt.Println(googleOauthConfig.ClientSecret)

	// Exchange the authorization code for an access token
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Token exchange error:", err) // Debug
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to exchange token",
		})
	}

	fmt.Println("Token exchanged successfully")

	// Fetch user info from Google
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch user info",
		})
	}
	defer resp.Body.Close()

	fmt.Println("User info fetched successfully")

	// Parse user info
	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse user info",
		})
	}

	fmt.Println("User info parsed successfully")

	// Extract relevant user details
	email := userInfo["email"].(string)
	name := userInfo["name"].(string)

	fmt.Println("Email:", email)
	fmt.Println("Name:", name)

	// Check if the user already exists in the database
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		// If the user doesn't exist, create a new user
		user = models.User{
			Name:  name,
			Email: email,
			// You can set a random password or leave it empty since this is a Google login
			Password: []byte(""), // No password for Google-authenticated users
		}

		// Insert the new user into the database
		if err := database.DB.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}
	}

	fmt.Println("User created/found successfully")

	// Generate a JWT token for the user
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})
	jwtToken, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	fmt.Println("JWT token generated successfully")

	// Rediriger vers le frontend avec le JWT dans l'URL
	frontendURL := "http://localhost:3000/auth/callback" // Remplacez par l'URL de votre frontend
	redirectURL := fmt.Sprintf("%s?token=%s", frontendURL, jwtToken)
	return c.Redirect(redirectURL, fiber.StatusFound)
}
