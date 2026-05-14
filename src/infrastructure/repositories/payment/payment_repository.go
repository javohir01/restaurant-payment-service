package payment

import (
	"context"

	"github.com/javohir01/restaurant-payment-service/src/domain/payment/models"
	"gorm.io/gorm"
)

type ReceiptRepo struct {
	db *gorm.DB
}

type TransactionRepo struct {
	db *gorm.DB
}

func NewReceiptRepo(db *gorm.DB) *ReceiptRepo {
	return &ReceiptRepo{db: db}
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (r *ReceiptRepo) SaveReceipt(ctx context.Context, receipt *models.Receipt) error {
	return r.db.WithContext(ctx).Save(receipt).Error
}

func (r *ReceiptRepo) GetReceipt(ctx context.Context, receiptID string) (*models.Receipt, error) {
	var receipt models.Receipt
	if err := r.db.WithContext(ctx).First(&receipt, "id = ?", receiptID).Error; err != nil {
		return nil, err
	}
	return &receipt, nil
}

func (r *ReceiptRepo) GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error) {
	var receipt models.Receipt
	if err := r.db.WithContext(ctx).First(&receipt, "order_id = ?", orderID).Error; err != nil {
		return nil, err
	}
	return &receipt, nil
}

func (r *TransactionRepo) SaveTx(ctx context.Context, tx *models.Transaction) error {
	return r.db.WithContext(ctx).Save(tx).Error
}

func (r *TransactionRepo) UpdateTx(ctx context.Context, tx *models.Transaction) error {
	return r.db.WithContext(ctx).Save(tx).Error
}

func (r *TransactionRepo) GetTx(ctx context.Context, txID string) (*models.Transaction, error) {
	var tx models.Transaction
	if err := r.db.WithContext(ctx).First(&tx, "id = ?", txID).Error; err != nil {
		return nil, err
	}
	return &tx, nil
}
