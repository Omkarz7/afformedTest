package routes

import (
	"affordmed/handlers"

	"github.com/gin-gonic/gin"
)

func OpenRoutes(open *gin.RouterGroup) {
	open.POST("/register", handlers.RegisterUser())
	open.POST("/login", handlers.Login())
}
