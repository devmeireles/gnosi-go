package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	gorm.Model
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	Username  string `json:"username" gorm:"unique;not null"`
	Biography string `json:"biography" gorm:""`
	Status    uint   `json:"status" gorm:"default:1"`
	Public    uint   `json:"public" gorm:"default:0"`
	TypeID    uint   `json:"type_id" gorm:"default:3"`
	AvatarID  uint   `json:"avatar_id" gorm:""`
	Language  string `json:"language" gorm:"default:'en'"`
	// Address  Address `json:"-"`
}

// UserLogin struct
type UserLogin struct {
	Username string
	Password string
}

// HashPassword hash the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // 14 is the cost for hashing the password.
	return string(bytes), err
}

// CheckPasswordHash checks if the password hash is able
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errors.New("password incorrect")
	}
	return nil
}

// BeforeSave sets a hashed password
func (u *User) BeforeSave() (err error) {
	password := strings.TrimSpace(u.Password)
	hashedpassword, err := HashPassword(password)
	if err != nil {
		return err
	}
	u.Password = string(hashedpassword)
	return nil
}
