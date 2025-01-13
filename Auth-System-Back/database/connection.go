// database/connection.go

package database

// Import GORM and PostgreSQL driver
import (
	"albus-auth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to PostgreSQL database
func ConnectDB() (*gorm.DB, error) {
	// PostgreSQL connection string format
	dsn := "host=localhost user=postgres password=password dbname=dbname port=5432 sslmode=disable"

	// Connect to PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	db.AutoMigrate(&models.User{})

	return db, nil
}
