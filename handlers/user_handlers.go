package handlers

import (
	"net/http"

	"github.com/Blaziant/go_balance_avito/database"
	"github.com/Blaziant/go_balance_avito/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"erroe": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "userBalance": user.Balance})
}
