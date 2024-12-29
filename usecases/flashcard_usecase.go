package usecases

import (
	"flashcard/models"
	"flashcard/repositories"
)

type FlashcardUsecase struct {
	Repository *repositories.FlashcardRepository
}

func NewFlashcardUsecase(repo *repositories.FlashcardRepository) *FlashcardUsecase {
	return &FlashcardUsecase{Repository: repo}
}

func (uc *FlashcardUsecase) CreateFlashcard(flashcard *models.Flashcard) error {
	return uc.Repository.Create(flashcard)
}

func (uc *FlashcardUsecase) GetFlashcardByID(id uint) (*models.Flashcard, error) {
	return uc.Repository.GetByID(id)
}

func (uc *FlashcardUsecase) UpdateFlashcard(flashcard *models.Flashcard) error {
	return uc.Repository.Update(flashcard)
}

func (uc *FlashcardUsecase) DeleteFlashcard(id uint) error {
	return uc.Repository.Delete(id)
}

func (uc *FlashcardUsecase) GetFlashcardsByDeckID(deckID uint) ([]models.Flashcard, error) {
	return uc.Repository.GetByDeckID(deckID)
}
