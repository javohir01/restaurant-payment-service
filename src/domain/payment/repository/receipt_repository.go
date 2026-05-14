package repository

import (
	"context"

	"github.com/javohir01/restaurant-payment-service/src/domain/payment/models"
)

type ReceiptRepository interface {
	SaveReceipt(ctx context.Context, receipt *models.Receipt) error
	GetReceipt(ctx context.Context, receiptID string) (*models.Receipt, error)
	GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error)
}
