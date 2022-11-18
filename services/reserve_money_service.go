package services

import (
	"github.com/Blaziant/go_balance_avito/errors"
	"github.com/Blaziant/go_balance_avito/models"
	"gorm.io/gorm"
)

func ReserveMoneyService(order_id uint, db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	var order models.Order
	if result := tx.First(&order, order_id); result.Error != nil {
		return result.Error
	}
	if order.ReserveAccountID != nil {
		return &errors.ReserveAccountUsed{Message: "Reserve account used"}
	}

	reserve_account := models.ReserveAccount{Order: order}
	if result := tx.Create(&reserve_account); result.Error != nil {
		return result.Error
	}

	var user models.User
	if result := tx.First(&user, order.UserID); result.Error != nil {
		return result.Error
	}
	if user.Balance < order.TotalPrice {
		return &errors.NotEnoughtMoneyError{Message: "User does not have enough money"}
	}
	user.Balance -= order.TotalPrice
	reserve_account.Balance += order.TotalPrice
	tx.Save(&user)
	tx.Save(&reserve_account)
	return tx.Commit().Error
}
