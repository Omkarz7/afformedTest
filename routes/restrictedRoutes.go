package routes

import "github.com/gin-gonic/gin"

func RestrictedRoutes(res *gin.RouterGroup) {
	res.POST("/edit-details")
	res.POST("/change-password")
}
