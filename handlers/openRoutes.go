package handlers

import (
	"affordmed/models"
	"affordmed/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var details models.UserDetials
		err := ctx.Bind(&details)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		models.RegisteredUsers = append(models.RegisteredUsers, details)

		ctx.JSON(http.StatusAccepted, gin.H{"message": "User Registered Successfully"})
	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var credentials models.Login
		err := ctx.Bind(&credentials)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ok := services.VerifyCredentails(credentials)
		if ok {
			//add a token to header
			ctx.JSON(http.StatusAccepted, gin.H{"message": "Credentials Verified"})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials"})
		}
	}
}
