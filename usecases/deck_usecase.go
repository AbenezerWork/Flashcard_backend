package usecases

import (
	"flashcard/models"
	"flashcard/repositories"
)

type DeckUsecase struct {
	Repository *repositories.DeckRepository
}

func NewDeckUsecase(repo *repositories.DeckRepository) *DeckUsecase {
	return &DeckUsecase{Repository: repo}
}

func (uc *DeckUsecase) CreateDeck(deck *models.Deck) error {
	return uc.Repository.Create(deck)
}

func (uc *DeckUsecase) GetDeckByID(id uint) (*models.Deck, error) {
	return uc.Repository.GetByID(id)
}

func (uc *DeckUsecase) GetDecksByUserID(userID uint) ([]models.Deck, error) {
	return uc.Repository.GetByUserID(userID)
}

func (uc *DeckUsecase) UpdateDeck(deck *models.Deck) error {
	return uc.Repository.Update(deck)
}

func (uc *DeckUsecase) DeleteDeck(id uint) error {
	return uc.Repository.Delete(id)
}
