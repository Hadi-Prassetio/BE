package models

import "time"

type Product struct {
	ID        int    `json:"id" gorm:"PRIMARY_KEY"`
	Title     string `json:"title" gorm:"type: varchar(255)"`
	Price     int    `json:"price" gorm:"type: varchar(255)"`
	Image     string `json:"image" gorm:"type: varchar(255)"`
	Cart      []Cart
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ProductResponse struct {
	Title int `json:"title"`
	Price int `json:"price"`
	Image int `json:"image"`
}

func (ProductResponse) TableName() string {
	return "users"
}
