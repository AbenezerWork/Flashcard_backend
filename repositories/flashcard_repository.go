package repositories

import (
	"flashcard/models"

	"gorm.io/gorm"
)

type FlashcardRepository struct {
	DB *gorm.DB
}

func NewFlashcardRepository(db *gorm.DB) *FlashcardRepository {
	return &FlashcardRepository{DB: db}
}

func (r *FlashcardRepository) Create(flashcard *models.Flashcard) error {
	return r.DB.Create(flashcard).Error
}

func (r *FlashcardRepository) GetByID(id uint) (*models.Flashcard, error) {
	var flashcard models.Flashcard
	err := r.DB.First(&flashcard, id).Error
	return &flashcard, err
}

func (r *FlashcardRepository) Update(flashcard *models.Flashcard) error {
	return r.DB.Save(flashcard).Error
}

func (r *FlashcardRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Flashcard{}, id).Error
}

func (r *FlashcardRepository) GetByDeckID(deckID uint) ([]models.Flashcard, error) {
	var flashcards []models.Flashcard
	err := r.DB.Where("deck_id = ?", deckID).Find(&flashcards).Error
	return flashcards, err
}
