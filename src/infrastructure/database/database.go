package database

import (
	"fmt"

	cardmodels "github.com/javohir01/restaurant-payment-service/src/domain/card/models"
	paymentmodels "github.com/javohir01/restaurant-payment-service/src/domain/payment/models"
	merchantmodels "github.com/javohir01/restaurant-payment-service/src/domain/merchant/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %w", err)
	}

	if err := db.AutoMigrate(
		&cardmodels.PaymentCard{},
		&merchantmodels.MerchantSetting{},
		&paymentmodels.Receipt{},
		&paymentmodels.Transaction{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}
