package models

import "time"

type MerchantSetting struct {
	EntityID  string    `json:"entity_id" gorm:"primaryKey"`
	CashboxID string    `json:"cashbox_id"`
	Enabled   bool      `json:"enabled"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
