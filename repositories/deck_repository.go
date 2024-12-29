package repositories

import (
	"flashcard/models"

	"gorm.io/gorm"
)

type DeckRepository struct {
	DB *gorm.DB
}

func NewDeckRepository(db *gorm.DB) *DeckRepository {
	return &DeckRepository{DB: db}
}

func (r *DeckRepository) Create(deck *models.Deck) error {
	return r.DB.Create(deck).Error
}

func (r *DeckRepository) GetByID(id uint) (*models.Deck, error) {
	var deck models.Deck
	err := r.DB.Preload("Flashcards").First(&deck, id).Error
	return &deck, err
}

func (r *DeckRepository) Update(deck *models.Deck) error {
	return r.DB.Save(deck).Error
}

func (r *DeckRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Deck{}, id).Error
}

func (r *DeckRepository) GetByUserID(userID uint) ([]models.Deck, error) {
	var decks []models.Deck
	err := r.DB.Where("user_id = ?", userID).Preload("Flashcards").Find(&decks).Error
	return decks, err
}
