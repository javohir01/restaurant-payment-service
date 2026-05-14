package models

import "time"

type Receipt struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	OrderID   string    `json:"order_id" gorm:"index:idx_order_receipt"`
	CardID    string    `json:"card_id"`
	Amount    int       `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
