package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Balance float64 `json:"balance"`
	Orders  []Order
}

type Favour struct {
	gorm.Model
	Name   string   `json:"name"`
	Price  float64  `json:"price"`
	Orders []*Order `gorm:"many2many:order_favours;"`
}

type Order struct {
	gorm.Model
	TotalPrice       float64
	Favours          []*Favour `gorm:"many2many:order_favours;"`
	UserID           uint    
	ReserveAccountID *uint
	Accepted         bool
}

type ReserveAccount struct {
	gorm.Model
	Balance float64 `json:"balance"`
	Order   Order
}
