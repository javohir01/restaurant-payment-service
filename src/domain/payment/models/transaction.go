package models

import "time"

type Transaction struct {
	ID           string            `json:"id" gorm:"primaryKey"`
	RestaurantID string            `json:"restaurant_id" gorm:"index:idx_restaurant_tx"`
	Currency     string            `json:"currency"`
	Status       string            `json:"status"`
	CreateTime   int64             `json:"create_time"`
	PayTime      int64             `json:"pay_time"`
	CancelTime   int64             `json:"cancel_time"`
	CardID       string            `json:"card_id"`
	Amount       int               `json:"amount"`
	Account      map[string]string `json:"account" gorm:"serializer:json"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}
