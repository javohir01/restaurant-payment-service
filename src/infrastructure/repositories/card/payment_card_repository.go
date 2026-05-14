package card

import (
	"context"

	"github.com/javohir01/restaurant-payment-service/src/domain/card/models"
	"gorm.io/gorm"
)

type PaymentCardRepo struct {
	db *gorm.DB
}

func NewPaymentCardRepo(db *gorm.DB) *PaymentCardRepo {
	return &PaymentCardRepo{db: db}
}

func (r *PaymentCardRepo) SaveCard(ctx context.Context, card *models.PaymentCard) error {
	return r.db.WithContext(ctx).Save(card).Error
}

func (r *PaymentCardRepo) DeleteCard(ctx context.Context, cardID string) error {
	return r.db.WithContext(ctx).Delete(&models.PaymentCard{}, "id = ?", cardID).Error
}

func (r *PaymentCardRepo) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {
	var card models.PaymentCard
	if err := r.db.WithContext(ctx).First(&card, "id = ?", cardID).Error; err != nil {
		return nil, err
	}
	return &card, nil
}
