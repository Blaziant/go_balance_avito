package services

import (
	"github.com/Blaziant/go_balance_avito/errors"
	"github.com/Blaziant/go_balance_avito/models"
	"gorm.io/gorm"
)

type TransferMoneySchema struct {
	FromId uint    `json:"from_id" binding:"required"`
	ToId   uint    `json:"to_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

func TranserMoneyService(db *gorm.DB, data *TransferMoneySchema) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var from_user models.User
	record := tx.First(&from_user, data.FromId)
	if err := record.Error; err != nil {
		return err
	}

	var to_user models.User
	record = tx.First(&to_user, data.ToId)
	if err := record.Error; err != nil {
		return err
	}
	if from_user.Balance < data.Amount {
		return &errors.NotEnoughtMoneyError{Message: "Not enought money to transfer."}
	}
	from_user.Balance -= data.Amount
	to_user.Balance += data.Amount
	tx.Save(&from_user)
	tx.Save(&to_user)

	return tx.Commit().Error
}
