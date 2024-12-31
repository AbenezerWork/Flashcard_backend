package routes

import (
	"flashcard/controllers"
	"flashcard/middleware"
	"flashcard/repositories"
	"flashcard/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	authController := controllers.NewAuthController(userUsecase)

	deckRepo := repositories.NewDeckRepository(db)
	deckUsecase := usecases.NewDeckUsecase(deckRepo)
	deckController := controllers.NewDeckController(deckUsecase)

	flashcardRepo := repositories.NewFlashcardRepository(db)
	flashcardUsecase := usecases.NewFlashcardUsecase(flashcardRepo)
	flashcardController := controllers.NewFlashcardController(flashcardUsecase, deckUsecase)

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
	r.GET("/validatetoken", authController.ValidateToken)
	api := r.Group("/api/v1")
	{
		api.Use(middleware.AuthMiddleware())

		api.POST("/flashcards", flashcardController.CreateFlashcard)
		api.GET("/flashcards/:id", flashcardController.GetFlashcardByID)
		api.PUT("/flashcards/:id", flashcardController.UpdateFlashcard)
		api.DELETE("/flashcards/:id", flashcardController.DeleteFlashcard)

		api.POST("/decks", deckController.CreateDeck)
		api.GET("/decks/:id", deckController.GetDeckByID)
		api.GET("/decks", deckController.GetDecksByUser)
		api.PUT("/decks/:id", deckController.UpdateDeck)
		api.DELETE("/decks/:id", deckController.DeleteDeck)
	}

	return r
}
