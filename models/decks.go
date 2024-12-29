package models

import (
	"gorm.io/gorm"
)

type Deck struct {
	gorm.Model
	UserID     uint        `json:"user_id"`
	Name       string      `gorm:"type:varchar(100);not null" json:"name"`
	Flashcards []Flashcard `gorm: "foreignKey:DeckID" json:"flashcards"`
}
