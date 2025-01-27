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
	"golang.org/x/oauth2/linkedin"
)

func LinkedInCallback(c *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env: %v", err)
	}

	var (
		linkedinOauthConfig = &oauth2.Config{
			RedirectURL:  "http://localhost:8000/auth/linkedin/callback",
			ClientID:     os.Getenv("LINKEDIN_CLIENT_ID"),
			ClientSecret: os.Getenv("LINKEDIN_CLIENT_SECRET"),
			Scopes:       []string{"openid", "profile", "email"}, // Scopes pour accéder au profil et à l'email
			Endpoint:     linkedin.Endpoint,
		}
	)

	// Récupérer le code et l'état des paramètres de requête
	code := c.Query("code")
	state := c.Query("state")

	fmt.Println("Code:", code)
	fmt.Println("State:", state)
	fmt.Println(linkedinOauthConfig.ClientID)
	fmt.Println(linkedinOauthConfig.ClientSecret)

	// Valider l'état (optionnel mais recommandé)
	if state == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing state parameter",
		})
	}

	// Échanger le code contre un jeton d'accès
	token, err := linkedinOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("Token exchange error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to exchange token",
		})
	}

	// Récupérer les informations de l'utilisateur
	client := linkedinOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://api.linkedin.com/v2/me?projection=(id,localizedFirstName,localizedLastName)")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch user info",
		})
	}
	defer resp.Body.Close()

	var profile struct {
		ID                 string `json:"id"`
		LocalizedFirstName string `json:"localizedFirstName"`
		LocalizedLastName  string `json:"localizedLastName"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse user info",
		})
	}

	/**	// Récupérer l'email de l'utilisateur
	emailResp, err := client.Get("https://api.linkedin.com/v2/emailAddress?q=members&projection=(elements*(handle~))")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch email",
		})
	}
	defer emailResp.Body.Close()

	var emailData struct {
		Elements []struct {
			Handle struct {
				EmailAddress string `json:"emailAddress"`
			} `json:"handle~"`
		} `json:"elements"`
	}
	if err := json.NewDecoder(emailResp.Body).Decode(&emailData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse email",
		})
	}

	fmt.Println("Email Response Status:", emailResp.Status)
	fmt.Println("Email Response Body:", emailData)

	// Vérifier si l'email est disponible
	if len(emailData.Elements) == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No email found in LinkedIn response",
		})
	}

	email := emailData.Elements[0].Handle.EmailAddress
	*/

	resp, err = client.Get("https://api.linkedin.com/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch user info",
		})
	}
	defer resp.Body.Close()

	var userInfo struct {
		Sub     string `json:"sub"`     // ID utilisateur
		Name    string `json:"name"`    // Nom complet
		Email   string `json:"email"`   // Adresse e-mail
		Picture string `json:"picture"` // URL de la photo de profil
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse user info",
		})
	}

	// Utiliser les informations de l'utilisateur
	email := userInfo.Email
	name := userInfo.Name

	// Vérifier si l'utilisateur existe déjà dans la base de données
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		// Créer un nouvel utilisateur s'il n'existe pas
		user = models.User{
			Name:     name, //profile.LocalizedFirstName + " " + profile.LocalizedLastName,
			Email:    email,
			Password: []byte(""), // Pas de mot de passe pour les utilisateurs LinkedIn
		}
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
	jwtToken, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// Rediriger vers le frontend avec le JWT
	frontendURL := "http://localhost:3000/auth/callback" // URL de votre frontend
	redirectURL := fmt.Sprintf("%s?token=%s", frontendURL, jwtToken)
	return c.Redirect(redirectURL, fiber.StatusFound)
}
