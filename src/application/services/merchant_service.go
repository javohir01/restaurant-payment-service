package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/javohir01/restaurant-payment-service/src/domain/merchant/models"
	"github.com/javohir01/restaurant-payment-service/src/domain/merchant/repository"
)

type MerchantService interface {
	CreateMerchantSetting(ctx context.Context, entityID, cashboxID string, enabled bool) (*models.MerchantSetting, error)
	UpdateMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error
	GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error)
}

type merchantService struct {
	repo repository.MerchantRepository
}

func NewMerchantService(repo repository.MerchantRepository) MerchantService {
	return &merchantService{repo: repo}
}

func (s *merchantService) CreateMerchantSetting(ctx context.Context, entityID, cashboxID string, enabled bool) (*models.MerchantSetting, error) {
	if entityID == "" {
		entityID = uuid.NewString()
	}
	if cashboxID == "" {
		return nil, errors.New("cashbox id is required")
	}

	setting := &models.MerchantSetting{
		EntityID:  entityID,
		CashboxID: cashboxID,
		Enabled:   enabled,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := s.repo.SaveMerchantSetting(ctx, setting); err != nil {
		return nil, err
	}
	return setting, nil
}

func (s *merchantService) UpdateMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error {
	if setting == nil {
		return errors.New("merchant setting is required")
	}
	if setting.EntityID == "" {
		return errors.New("entity id is required")
	}
	setting.UpdatedAt = time.Now().UTC()
	return s.repo.UpdateMerchantSetting(ctx, setting)
}

func (s *merchantService) GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error) {
	if entityID == "" {
		return nil, errors.New("entity id is required")
	}
	return s.repo.GetMerchantSetting(ctx, entityID)
}
