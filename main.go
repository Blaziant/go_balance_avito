package main

import (
	"github.com/Blaziant/go_balance_avito/database"
	"github.com/Blaziant/go_balance_avito/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	config := database.Config{}
	config.FillConfig()
	database.Connect(config.DatabaseUrl())
	database.Migrate()

	router := initRouters()
	router.Run(":8000")
}

func initRouters() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.GET("/:id/balance", handlers.GetBalance)
			user.POST("/debit_balance", handlers.DebitBalance)
			user.POST("/replenish_balance", handlers.ReplenishBalance)
			user.POST("/transfer_money", handlers.TransferMoney)
			user.POST("", handlers.CreateUser)
		}

		favour := api.Group("/favour")
		{
			favour.GET("/:id", handlers.GetFavour)
			favour.POST("", handlers.CreateFavour)
		}

		order := api.Group("/order")
		{
			order.POST("/accept", handlers.AcceptOrder)
			order.POST("/reserve_money", handlers.ReserveMoney)
			order.POST("", handlers.CreateOrder)
		}
	}
	return router
}
