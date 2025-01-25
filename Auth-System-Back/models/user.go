// models/user.go
package models

type User struct {
	ID       string `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password []byte `json:"-"`
}
