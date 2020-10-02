package handlers

import (
	"affordmed/models"
	"affordmed/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EditUserDetials() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var details models.UserDetials
		err := ctx.Bind(&details)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ok := services.EditUserDetials(details)
		if ok {
			//add a token to header
			ctx.JSON(http.StatusAccepted, gin.H{"message": "User Details Updated Succesfully"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Unable to find user"})
		}

	}
}
