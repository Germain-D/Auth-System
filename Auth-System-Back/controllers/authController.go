// controllers/authController.go

package controllers

import (
	"albus-auth/database"
	"albus-auth/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "secret"

// Hello returns a simple "Hello world!!" message
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello world!!")
}

// controllers/authController.go

func Register(c *fiber.Ctx) error {
	fmt.Println("Received a registration request")

	// Parse request body
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	// Check if the email already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", data["email"]).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already exists",
		})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Create new user
	user := models.User{
		Name:     data["pseudo"],
		Email:    data["email"],
		Password: hashedPassword,
	}

	// Insert user into database
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
	})

}

func Login(c *fiber.Ctx) error {
	fmt.Println("Received a Login request")

	// Parse request body
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	// Check if user exists
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.ID == "" {
		fmt.Println("User not found")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	fmt.Println("User found")

	// Compare passwords
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil {
		fmt.Println("Invalid Password:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	fmt.Println("Password is correct")

	// Generate JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	})
	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// Set JWT token in cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // Expires in 24 hours
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(&cookie)

	// Authentication successful, return success response
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Login successful",
	})
}

func SimpleLogin(c *fiber.Ctx) error {
	// Hardcoded credentials
	const (
		knownUsername = "admin"
		knownPassword = "password123"
	)

	// Get credentials from request body
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Check if credentials match
	if data["username"] != knownUsername || data["password"] != knownPassword {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Generate JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": knownUsername,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// Set JWT token in cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Simple login successful",
	})
}

func User(c *fiber.Ctx) error {
	fmt.Println("Request to get user...")

	// Retrieve JWT token from cookie
	cookie := c.Cookies("jwt")

	// Parse JWT token with claims
	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	// Handle token parsing errors
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	// Extract claims from token
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse claims",
		})
	}

	// Extract user ID from claims
	id, _ := strconv.Atoi((*claims)["sub"].(string))
	user := models.User{ID: strconv.Itoa(id)}

	// Query user from database using ID
	database.DB.Where("id =?", id).First(&user)

	// Return user details as JSON response
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	fmt.Println("Received a logout request")

	// Clear JWT token by setting an empty value and expired time in the cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Expired 1 hour ago
		HTTPOnly: true,
		Secure:   true,
	}
	c.Cookie(&cookie)

	// Return success response indicating logout was successful
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Logout successful",
	})
}
