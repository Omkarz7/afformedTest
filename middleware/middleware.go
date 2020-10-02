package middleware

import (
	"affordmed/routes"

	"github.com/gin-gonic/gin"
)

func InitMiddleWare(router *gin.Engine) {
	open := router.Group("/o")
	routes.OpenRoutes(open)

	restricted := router.Group("/r")
	routes.RestrictedRoutes(restricted)
	// add a middleware to verify token for retricted routes
}
