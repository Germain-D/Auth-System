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
	"golang.org/x/oauth2/github"
)

func GitHubCallback(c *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env: %v", err)
	}

	var (
		githubOauthConfig = &oauth2.Config{
			RedirectURL:  os.Getenv("GITHUB_REDIRECT_URI"),
			ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
			Scopes:       []string{"user:email"}, // Scopes pour accéder à l'e-mail de l'utilisateur
			Endpoint:     github.Endpoint,
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
	token, err := githubOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to exchange token",
		})
	}

	// Récupérer les informations de l'utilisateur
	client := githubOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch user info",
		})
	}
	defer resp.Body.Close()

	var profile struct {
		ID    int    `json:"id"`
		Login string `json:"login"` // Nom d'utilisateur GitHub
		Name  string `json:"name"`  // Nom complet
		Email string `json:"email"` // E-mail (peut être null)
	}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse user info",
		})
	}

	// Si l'e-mail n'est pas disponible dans la réponse, le récupérer via l'API GitHub
	if profile.Email == "" {
		emailResp, err := client.Get("https://api.github.com/user/emails")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch email",
			})
		}
		defer emailResp.Body.Close()

		var emails []struct {
			Email   string `json:"email"`
			Primary bool   `json:"primary"`
		}
		if err := json.NewDecoder(emailResp.Body).Decode(&emails); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to parse emails",
			})
		}

		// Trouver l'e-mail principal
		for _, email := range emails {
			if email.Primary {
				profile.Email = email.Email
				break
			}
		}
	}

	// Vérifier si l'utilisateur existe déjà dans la base de données
	var user models.User
	if err := database.DB.Where("email = ?", profile.Email).First(&user).Error; err != nil {
		// Créer un nouvel utilisateur s'il n'existe pas
		user = models.User{
			Name:     profile.Name,
			Email:    profile.Email,
			Password: []byte(""), // Pas de mot de passe pour les utilisateurs GitHub
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
