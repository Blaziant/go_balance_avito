package handlers

import (
	"net/http"
	"time"

	"github.com/Blaziant/go_balance_avito/database"
	"github.com/Blaziant/go_balance_avito/models"
	"github.com/Blaziant/go_balance_avito/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type CreateOrderSchema struct {
	FavoursID []uint `json:"favours_id" binding:"required"`
	UserID    uint   `json:"user_id" binding:"required"`
}

// Создает нoвый заказ
func CreateOrder(context *gin.Context) {
	var data CreateOrderSchema
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var favours []*models.Favour
	if result := database.Instance.Where(data.FavoursID).Find(&favours); result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		context.Abort()
		return
	}
	var total_price float64
	for _, favour := range favours {
		total_price += favour.Price
	}

	order := models.Order{TotalPrice: total_price, UserID: data.UserID, Favours: favours}
	if record := database.Instance.Create(&order); record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusCreated, gin.H{"order_id": order.ID, "total_price": total_price, "user_id": order.UserID, "accepted": order.Accepted})
}

type OrderIDSchema struct {
	OrderID uint `json:"order_id" binding:"required"`
}

func AcceptOrder(context *gin.Context) {
	var data OrderIDSchema
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := services.AcceptOrderService(data.OrderID, database.Instance); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"accepted": true})
}

func ReserveMoney(context *gin.Context) {
	var data OrderIDSchema
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := services.ReserveMoneyService(data.OrderID, database.Instance); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "ok"})
}

type BookkeepingReportSchema struct {
	BeginAt string `json:"begin" binding:"required"`
	EndAt   string `json:"end" binding:"required"`
}

func BookkeepingReport(context *gin.Context) {
	var period BookkeepingReportSchema
	if err := context.ShouldBindJSON(&period); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	begin_date, err := time.Parse("2006-01-02", period.BeginAt)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	end_date, err := time.Parse("2006-01-02", period.EndAt)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	var orders []*models.Order
	if result := database.Instance.Where("accepted = ?", true).Where("created_at between ? and ?", begin_date, end_date).Preload(clause.Associations).Find(&orders); result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"orders": orders})
}
