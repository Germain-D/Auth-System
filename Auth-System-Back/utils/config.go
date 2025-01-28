package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all the environment variables.
type Config struct {
	// Google OAuth
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURI  string

	// LinkedIn OAuth
	LinkedInClientID     string
	LinkedInClientSecret string
	LinkedInRedirectURI  string

	// GitHub OAuth
	GitHubClientID     string
	GitHubClientSecret string
	GitHubRedirectURI  string

	// Facebook OAuth
	FacebookClientID     string
	FacebookClientSecret string
	FacebookRedirectURI  string

	// Database
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string

	// Frontend
	FrontendURL string

	// Server
	ServerPort string

	// Secret Key
	SecretKey string

	// Simple Auth
	SimpleUsername string
	SimplePassword string

	//log level
	LogLevel string
}

func LoadConfig() (*Config, error) {
	sugar := SugaredLogger

	if err := godotenv.Load(); err != nil {
		sugar.Warnw("Warning: .env file not found",
			"error", err,
		)
	}

	return &Config{
		// Google OAuth
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirectURI:  getEnv("GOOGLE_REDIRECT_URI", ""),

		// LinkedIn OAuth
		LinkedInClientID:     getEnv("LINKEDIN_CLIENT_ID", ""),
		LinkedInClientSecret: getEnv("LINKEDIN_CLIENT_SECRET", ""),
		LinkedInRedirectURI:  getEnv("LINKEDIN_REDIRECT_URI", ""),

		// GitHub OAuth
		GitHubClientID:     getEnv("GITHUB_CLIENT_ID", ""),
		GitHubClientSecret: getEnv("GITHUB_CLIENT_SECRET", ""),
		GitHubRedirectURI:  getEnv("GITHUB_REDIRECT_URI", ""),

		// Facebook OAuth
		FacebookClientID:     getEnv("FACEBOOK_CLIENT_ID", ""),
		FacebookClientSecret: getEnv("FACEBOOK_CLIENT_SECRET", ""),
		FacebookRedirectURI:  getEnv("FACEBOOK_REDIRECT_URI", ""),

		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "auth-system"),
		DBPort:     getEnv("DB_PORT", "5432"),

		// Frontend
		FrontendURL: getEnv("FRONTEND_URL", "http://localhost:3000"),

		// Server
		ServerPort: getEnv("SERVER_PORT", ":8000"),

		// Secret Key
		SecretKey: getEnv("SECRET_KEY", "secret"),

		// Simple Auth
		SimpleUsername: getEnv("SIMPLE_USERNAME", "admin"),
		SimplePassword: getEnv("SIMPLE_PASSWORD", "admin"),

		//log level
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		SugaredLogger.Debugw("Using default value for config",
			"key", key,
			"default", defaultValue,
		)
		return defaultValue
	}
	return value
}

// getEnvAsInt retrieves an environment variable as an integer or returns a default value.
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Invalid value for %s, using default: %d\n", key, defaultValue)
		return defaultValue
	}
	return value
}
