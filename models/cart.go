package models

import "time"

type Cart struct {
	ID             int       `json:"id" gorm:"PRIMARY_KEY"`
	Product_ID     int       `json:"product_id"`
	Transaction_ID int       `json:"transaction_id"`
	Topping        []Topping `gorm:"many2many:cart_toppings;"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}
