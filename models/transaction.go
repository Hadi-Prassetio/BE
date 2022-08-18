package models

import "time"

type Transaction struct {
	ID        int `json:"id" gorm:"PRIMARY_KEY"`
	UserID    int `json:"user_id"`
	Cart      []Cart
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type TransactionResponse struct {
	UserID int                  `json:"-"`
	User   UsersProfileResponse `json:"user"`
}

func (TransactionResponse) TableName() string {
	return "profiles"
}
