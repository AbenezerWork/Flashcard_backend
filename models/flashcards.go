package models

import (
	"gorm.io/gorm"
)

type Flashcard struct {
	gorm.Model
	Question string `gorm:"type:text;not null" json:"question"`
	Answer   string `gorm:"type:text;not null" json:"answer"`
	Category string `gorm:"type:varchar(100)" json:"category"`
	DeckID   uint   `json:"deckid`
}
