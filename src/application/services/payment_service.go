package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/javohir01/restaurant-payment-service/src/domain/payment/models"
	receiptrepo "github.com/javohir01/restaurant-payment-service/src/domain/payment/repository"
	"github.com/javohir01/restaurant-payment-service/src/infrastructure/payments"
)

type PaymentService interface {
	CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error)
	PayReceipt(ctx context.Context, receiptID string) error
	GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error)
	SaveTransaction(ctx context.Context, tx *models.Transaction) error
	UpdateTransaction(ctx context.Context, tx *models.Transaction) error
	GetTransaction(ctx context.Context, txID string) (*models.Transaction, error)
}

type paymentService struct {
	receiptRepo      receiptrepo.ReceiptRepository
	transactionRepo  receiptrepo.TransactionRepository
	paymentGateway   payments.PaymentGateway
}

func NewPaymentService(
	receiptRepo receiptrepo.ReceiptRepository,
	transactionRepo receiptrepo.TransactionRepository,
	paymentGateway payments.PaymentGateway,
) PaymentService {
	return &paymentService{
		receiptRepo:     receiptRepo,
		transactionRepo: transactionRepo,
		paymentGateway:  paymentGateway,
	}
}

func (s *paymentService) CreateReceipt(ctx context.Context, orderID, cardID string, amount int) (*models.Receipt, error) {
	if orderID == "" {
		return nil, errors.New("order id is required")
	}
	if cardID == "" {
		return nil, errors.New("card id is required")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	receipt := &models.Receipt{
		ID:        uuid.NewString(),
		OrderID:   orderID,
		CardID:    cardID,
		Amount:    amount,
		Status:    "pending",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := s.receiptRepo.SaveReceipt(ctx, receipt); err != nil {
		return nil, err
	}

	return receipt, nil
}

func (s *paymentService) PayReceipt(ctx context.Context, receiptID string) error {
	if receiptID == "" {
		return errors.New("receipt id is required")
	}

	receipt, err := s.receiptRepo.GetReceipt(ctx, receiptID)
	if err != nil {
		return err
	}
	if receipt.Status == "paid" {
		return errors.New("receipt already paid")
	}

	if err := s.paymentGateway.ConfirmPayment(ctx, receipt); err != nil {
		return err
	}

	receipt.Status = "paid"
	receipt.UpdatedAt = time.Now().UTC()
	return s.receiptRepo.SaveReceipt(ctx, receipt)
}

func (s *paymentService) GetReceiptByOrder(ctx context.Context, orderID string) (*models.Receipt, error) {
	if orderID == "" {
		return nil, errors.New("order id is required")
	}
	return s.receiptRepo.GetReceiptByOrder(ctx, orderID)
}

func (s *paymentService) SaveTransaction(ctx context.Context, tx *models.Transaction) error {
	if tx == nil {
		return errors.New("transaction is required")
	}
	if tx.ID == "" {
		tx.ID = uuid.NewString()
	}
	if tx.RestaurantID == "" {
		return errors.New("restaurant id is required")
	}
	if tx.Amount <= 0 {
		return errors.New("transaction amount must be greater than zero")
	}
	if tx.Currency == "" {
		return errors.New("currency is required")
	}

	tx.CreatedAt = time.Now().UTC()
	tx.UpdatedAt = time.Now().UTC()

	return s.transactionRepo.SaveTx(ctx, tx)
}

func (s *paymentService) UpdateTransaction(ctx context.Context, tx *models.Transaction) error {
	if tx == nil {
		return errors.New("transaction is required")
	}
	if tx.ID == "" {
		return errors.New("transaction id is required")
	}

	tx.UpdatedAt = time.Now().UTC()
	return s.transactionRepo.UpdateTx(ctx, tx)
}

func (s *paymentService) GetTransaction(ctx context.Context, txID string) (*models.Transaction, error) {
	if txID == "" {
		return nil, errors.New("transaction id is required")
	}
	return s.transactionRepo.GetTx(ctx, txID)
}
