package main

import (
	"log"
	"net/http"

	"github.com/javohir01/restaurant-payment-service/src/application/services"
	"github.com/javohir01/restaurant-payment-service/src/infrastructure/database"
	"github.com/javohir01/restaurant-payment-service/src/infrastructure/payments"
	cardRepo "github.com/javohir01/restaurant-payment-service/src/infrastructure/repositories/card"
	merchantRepo "github.com/javohir01/restaurant-payment-service/src/infrastructure/repositories/merchant"
	paymentRepo "github.com/javohir01/restaurant-payment-service/src/infrastructure/repositories/payment"
	api "github.com/javohir01/restaurant-payment-service/src/interfaces/http"
)

func main() {
	db, err := database.NewDatabase("restaurant_payment.db")
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	cardRepository := cardRepo.NewPaymentCardRepo(db)
	merchantRepository := merchantRepo.NewMerchantRepo(db)
	receiptRepository := paymentRepo.NewReceiptRepo(db)
	transactionRepository := paymentRepo.NewTransactionRepo(db)

	paymentGateway := payments.NewPaymeGateway()

	cardService := services.NewPaymentCardService(cardRepository)
	merchantService := services.NewMerchantService(merchantRepository)
	paymentService := services.NewPaymentService(receiptRepository, transactionRepository, paymentGateway)

	api.RegisterRoutes(http.DefaultServeMux, cardService, merchantService, paymentService)

	log.Println("Starting restaurant-payment-service on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
