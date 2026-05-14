package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/javohir01/restaurant-payment-service/src/domain/card/models"
	"github.com/javohir01/restaurant-payment-service/src/domain/card/repository"
)

type PaymentCardService interface {
	SaveCard(ctx context.Context, cardID, cardToken string) (*models.PaymentCard, error)
	DeleteCard(ctx context.Context, cardID string) error
	GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error)
}

type paymentCardService struct {
	repo repository.PaymentCardRepository
}

func NewPaymentCardService(repo repository.PaymentCardRepository) PaymentCardService {
	return &paymentCardService{repo: repo}
}

func (s *paymentCardService) SaveCard(ctx context.Context, cardID, cardToken string) (*models.PaymentCard, error) {
	if cardToken == "" {
		return nil, errors.New("card token is required")
	}

	if cardID == "" {
		cardID = uuid.NewString()
	}

	card := &models.PaymentCard{
		ID:        cardID,
		CardToken: cardToken,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.SaveCard(ctx, card); err != nil {
		return nil, err
	}

	return card, nil
}

func (s *paymentCardService) DeleteCard(ctx context.Context, cardID string) error {
	if cardID == "" {
		return errors.New("card id is required")
	}
	return s.repo.DeleteCard(ctx, cardID)
}

func (s *paymentCardService) GetCard(ctx context.Context, cardID string) (*models.PaymentCard, error) {
	if cardID == "" {
		return nil, errors.New("card id is required")
	}
	return s.repo.GetCard(ctx, cardID)
}
