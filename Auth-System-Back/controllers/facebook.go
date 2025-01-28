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
	"golang.org/x/oauth2/facebook"
)

func FacebookCallback(c *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env: %v", err)
	}

	var (
		facebookOauthConfig = &oauth2.Config{
			RedirectURL:  os.Getenv("FACEBOOK_REDIRECT_URI"),
			ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
			ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
			Scopes:       []string{"email"}, // Scopes pour accéder à l'e-mail de l'utilisateur
			Endpoint:     facebook.Endpoint,
		}
	)

	// Récupérer le code et l'état des paramètres de requête
	code := c.Query("code")
	state := c.Query("state")

	// Valider l'état (optionnel mais recommandé)
	if state == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing state parameter",
		})
	}

	// Échanger le code contre un jeton d'accès
	token, err := facebookOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to exchange token",
		})
	}

	// Récupérer les informations de l'utilisateur
	client := facebookOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://graph.facebook.com/v12.0/me?fields=id,name,email")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch user info",
		})
	}
	defer resp.Body.Close()

	var profile struct {
		ID    string `json:"id"`
		Name  string `json:"name"`  // Nom complet
		Email string `json:"email"` // E-mail
	}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse user info",
		})
	}

	// Vérifier si l'utilisateur existe déjà dans la base de données
	var user models.User
	if err := database.DB.Where("email = ?", profile.Email).First(&user).Error; err != nil {
		// Créer un nouvel utilisateur s'il n'existe pas
		user = models.User{
			Name:     profile.Name,
			Email:    profile.Email,
			Password: []byte(""), // Pas de mot de passe pour les utilisateurs Facebook
		}

		// Insérer le nouvel utilisateur dans la base de données
		if err := database.DB.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}
	}

	// Générer un JWT
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expire dans 24 heures
	})
	jwtToken, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// Rediriger vers le frontend avec le JWT
	frontendURL := os.Getenv("FRONTEND_URL") + "/auth/callback"
	redirectURL := fmt.Sprintf("%s?token=%s", frontendURL, jwtToken)
	return c.Redirect(redirectURL, fiber.StatusFound)
}
