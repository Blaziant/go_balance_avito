package handlers

import (
	"net/http"

	"github.com/Blaziant/go_balance_avito/database"
	"github.com/Blaziant/go_balance_avito/models"
	"github.com/gin-gonic/gin"
)

func CreateFavour(context *gin.Context) {
	var favour models.Favour
	if err := context.ShouldBindJSON(&favour); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"erroe": err.Error()})
		context.Abort()
		return
	}
	if record := database.Instance.Create(&favour); record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	if favour.Price < 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Negative favour price"})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"favourId": favour.ID, "FavourName": favour.Name, "FavourPrice": favour.Price})
}

func GetFavour(context *gin.Context) {
	var favour models.Favour
	if record := database.Instance.First(&favour, context.Param("id")); record.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"favour": favour})
}
