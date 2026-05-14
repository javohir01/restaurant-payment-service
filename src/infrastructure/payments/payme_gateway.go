package payments

import (
	"context"
	"fmt"

	"github.com/javohir01/restaurant-payment-service/src/domain/payment/models"
)

type PaymeGateway struct{}

type PaymentGateway interface {
	ConfirmPayment(ctx context.Context, receipt *models.Receipt) error
}

func NewPaymeGateway() *PaymeGateway {
	return &PaymeGateway{}
}

func (g *PaymeGateway) ConfirmPayment(ctx context.Context, receipt *models.Receipt) error {
	if receipt == nil {
		return fmt.Errorf("receipt is required")
	}

	if receipt.Amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}

	// Simulated Payme API integration.
	return nil
}
