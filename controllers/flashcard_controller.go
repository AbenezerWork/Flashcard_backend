package controllers

import (
	"net/http"
	"strconv"

	"flashcard/models"
	"flashcard/usecases"

	"github.com/gin-gonic/gin"
)

type FlashcardController struct {
	Usecase     *usecases.FlashcardUsecase
	DeckUsecase *usecases.DeckUsecase
}

func NewFlashcardController(uc *usecases.FlashcardUsecase, du *usecases.DeckUsecase) *FlashcardController {
	return &FlashcardController{Usecase: uc}
}

func (ctrl *FlashcardController) CreateFlashcard(c *gin.Context) {
	var flashcard models.Flashcard
	if err := c.ShouldBindJSON(&flashcard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Usecase.CreateFlashcard(&flashcard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, flashcard)
}

func (ctrl *FlashcardController) GetFlashcardByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	flashcard, err := ctrl.Usecase.GetFlashcardByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flashcard not found"})
		return
	}
	c.JSON(http.StatusOK, flashcard)
}

func (ctrl *FlashcardController) UpdateFlashcard(c *gin.Context) {
	var flashcard models.Flashcard
	if err := c.ShouldBindJSON(&flashcard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Usecase.UpdateFlashcard(&flashcard); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, flashcard)
}

func (ctrl *FlashcardController) DeleteFlashcard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := ctrl.Usecase.DeleteFlashcard(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Flashcard deleted"})
}
