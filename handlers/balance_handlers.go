package handlers

import (
	"net/http"

	"github.com/Blaziant/go_balance_avito/database"
	"github.com/Blaziant/go_balance_avito/models"
	"github.com/Blaziant/go_balance_avito/services"
	"github.com/gin-gonic/gin"
)

type ChangeBalanceSchema struct {
	Amount float64 `json:"amount" binding:"required"`
	UserId uint    `json:"user_id" binding:"required"`
}

func ReplenishBalance(context *gin.Context) {
	var data ChangeBalanceSchema
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var user models.User
	record := database.Instance.First(&user, data.UserId)
	if record.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	user.Balance += data.Amount
	database.Instance.Save(&user)
	context.JSON(http.StatusOK, gin.H{"new balance": user.Balance})
}

type GetBalanceSchema struct {
	UserId uint `json:"user_id" binding:"required"`
}

func GetBalance(context *gin.Context) {
	var user models.User
	record := database.Instance.First(&user, context.Param("id"))
	if record.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"balance": user.Balance})

}

func DebitBalance(context *gin.Context) {
	var data ChangeBalanceSchema
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var user models.User
	record := database.Instance.First(&user, data.UserId)
	if record.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	if user.Balance < data.Amount {
		context.JSON(http.StatusNotFound, gin.H{"error": "User balance is too low"})
		context.Abort()
		return
	}

	user.Balance -= data.Amount
	database.Instance.Save(&user)
	context.JSON(http.StatusOK, gin.H{"balance": user.Balance})
}

func TransferMoney(context *gin.Context) {
	var data services.TransferMoneySchema
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := services.TranserMoneyService(database.Instance, &data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "ok"})
}
