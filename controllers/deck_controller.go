package controllers

import (
	"net/http"
	"strconv"

	"flashcard/models"
	"flashcard/usecases"

	"github.com/gin-gonic/gin"
)

type DeckController struct {
	Usecase *usecases.DeckUsecase
}

func NewDeckController(uc *usecases.DeckUsecase) *DeckController {
	return &DeckController{Usecase: uc}
}

func (ctrl *DeckController) CreateDeck(c *gin.Context) {
	var deck models.Deck
	if err := c.ShouldBindJSON(&deck); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deck.UserID = c.MustGet("userID").(uint)
	if err := ctrl.Usecase.CreateDeck(&deck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deck)
}

func (ctrl *DeckController) GetDeckByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	deck, err := ctrl.Usecase.GetDeckByID(uint(id))
	if err != nil || deck.UserID != c.MustGet("userID").(uint) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Deck not found"})
		return
	}
	c.JSON(http.StatusOK, deck)
}

func (ctrl *DeckController) GetDecksByUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	decks, err := ctrl.Usecase.GetDecksByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, decks)
}

func (ctrl *DeckController) UpdateDeck(c *gin.Context) {
	var deck models.Deck
	if err := c.ShouldBindJSON(&deck); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Usecase.UpdateDeck(&deck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deck)
}

func (ctrl *DeckController) DeleteDeck(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := ctrl.Usecase.DeleteDeck(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deck deleted"})
}
