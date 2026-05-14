package merchant

import (
	"context"

	"github.com/javohir01/restaurant-payment-service/src/domain/merchant/models"
	"gorm.io/gorm"
)

type MerchantRepo struct {
	db *gorm.DB
}

func NewMerchantRepo(db *gorm.DB) *MerchantRepo {
	return &MerchantRepo{db: db}
}

func (r *MerchantRepo) SaveMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error {
	return r.db.WithContext(ctx).Save(setting).Error
}

func (r *MerchantRepo) UpdateMerchantSetting(ctx context.Context, setting *models.MerchantSetting) error {
	return r.db.WithContext(ctx).Save(setting).Error
}

func (r *MerchantRepo) GetMerchantSetting(ctx context.Context, entityID string) (*models.MerchantSetting, error) {
	var setting models.MerchantSetting
	if err := r.db.WithContext(ctx).First(&setting, "entity_id = ?", entityID).Error; err != nil {
		return nil, err
	}
	return &setting, nil
}
