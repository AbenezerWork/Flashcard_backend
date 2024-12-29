package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Decks    []Deck `gorm:"foreignKey:UserID" json:"decks"`
}

func CheckPassword(providedPassword, storedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
}
