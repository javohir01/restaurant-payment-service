package models

import "time"

type PaymentCard struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	CardToken string    `json:"card_token"`
	CreatedAt time.Time `json:"created_at"`
}
