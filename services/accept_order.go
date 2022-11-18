package services

import (
	"github.com/Blaziant/go_balance_avito/errors"
	"github.com/Blaziant/go_balance_avito/models"
	"gorm.io/gorm"
)

func AcceptOrderService(order_id uint, db *gorm.DB) error {
	var order models.Order
	if result := db.First(&order, order_id); result.Error != nil {
		return result.Error
	}

	var reserve_account models.ReserveAccount
	if result := db.First(&reserve_account, order.ReserveAccountID); result.Error != nil {
		return result.Error
	}

	if order.Accepted == true {
		return &errors.AcceptedOrder{Message: "Order already accepted"}
	}

	if reserve_account.Balance < order.TotalPrice {
		return &errors.NotEnoughtMoneyError{Message: "Reserve account does not have enough money"}
	}

	order.Accepted = true
	reserve_account.Balance -= order.TotalPrice

	db.Save(&order)
	db.Save(&reserve_account)
	return nil
}
