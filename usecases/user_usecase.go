package usecases

import (
	"flashcard/models"
	"flashcard/repositories"
)

type UserUsecase struct {
	Repository *repositories.UserRepository
}

func NewUserUsecase(repo *repositories.UserRepository) *UserUsecase {
	return &UserUsecase{Repository: repo}
}

func (uc *UserUsecase) CreateUser(flashcard *models.User) error {
	return uc.Repository.Create(flashcard)
}

func (uc *UserUsecase) GetUserByID(id uint) (*models.User, error) {
	return uc.Repository.GetByID(id)
}

func (uc *UserUsecase) GetUserByEmail(email string) (*models.User, error) {
	return uc.Repository.GetByEmail(email)
}

func (uc *UserUsecase) UpdateUser(flashcard *models.User) error {
	return uc.Repository.Update(flashcard)
}

func (uc *UserUsecase) DeleteUser(id uint) error {
	return uc.Repository.Delete(id)
}
