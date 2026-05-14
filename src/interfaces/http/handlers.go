package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/javohir01/restaurant-payment-service/src/application/services"
	"github.com/javohir01/restaurant-payment-service/src/domain/payment/models"
)

type createCardRequest struct {
	CardID    string `json:"card_id"`
	CardToken string `json:"card_token"`
}

type createMerchantRequest struct {
	EntityID  string `json:"entity_id"`
	CashboxID string `json:"cashbox_id"`
	Enabled   bool   `json:"enabled"`
}

type updateMerchantRequest struct {
	CashboxID string `json:"cashbox_id"`
	Enabled   bool   `json:"enabled"`
}

type createReceiptRequest struct {
	OrderID string `json:"order_id"`
	CardID  string `json:"card_id"`
	Amount  int    `json:"amount"`
}

type createTransactionRequest struct {
	ID           string            `json:"id"`
	RestaurantID string            `json:"restaurant_id"`
	Currency     string            `json:"currency"`
	Status       string            `json:"status"`
	CreateTime   int64             `json:"create_time"`
	PayTime      int64             `json:"pay_time"`
	CancelTime   int64             `json:"cancel_time"`
	CardID       string            `json:"card_id"`
	Amount       int               `json:"amount"`
	Account      map[string]string `json:"account"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func RegisterRoutes(router *http.ServeMux,
	cardService services.PaymentCardService,
	merchantService services.MerchantService,
	paymentService services.PaymentService,
) {
	router.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusOK, map[string]interface{}{"status": "ok", "time": time.Now().UTC()})
	})

	router.HandleFunc("/api/v1/cards", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var req createCardRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				respondError(w, http.StatusBadRequest, err)
				return
			}
			if req.CardToken == "" {
				respondError(w, http.StatusBadRequest, errors.New("card_token is required"))
				return
			}
			card, err := cardService.SaveCard(r.Context(), req.CardID, req.CardToken)
			if err != nil {
				respondError(w, http.StatusInternalServerError, err)
				return
			}
			respondJSON(w, http.StatusCreated, card)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/api/v1/cards/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/api/v1/cards/")
		if id == "" {
			respondError(w, http.StatusBadRequest, errors.New("card id is required"))
			return
		}
		switch r.Method {
		case http.MethodGet:
			card, err := cardService.GetCard(r.Context(), id)
			if err != nil {
				respondError(w, http.StatusNotFound, err)
				return
			}
			respondJSON(w, http.StatusOK, card)
		case http.MethodDelete:
			if err := cardService.DeleteCard(r.Context(), id); err != nil {
				respondError(w, http.StatusInternalServerError, err)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/api/v1/merchants", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var req createMerchantRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				respondError(w, http.StatusBadRequest, err)
				return
			}
			if req.CashboxID == "" {
				respondError(w, http.StatusBadRequest, errors.New("cashbox_id is required"))
				return
			}
			setting, err := merchantService.CreateMerchantSetting(r.Context(), req.EntityID, req.CashboxID, req.Enabled)
			if err != nil {
				respondError(w, http.StatusInternalServerError, err)
				return
			}
			respondJSON(w, http.StatusCreated, setting)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/api/v1/merchants/", func(w http.ResponseWriter, r *http.Request) {
		entityID := strings.TrimPrefix(r.URL.Path, "/api/v1/merchants/")
		if entityID == "" {
			respondError(w, http.StatusBadRequest, errors.New("entity id is required"))
			return
		}
		switch r.Method {
		case http.MethodGet:
			setting, err := merchantService.GetMerchantSetting(r.Context(), entityID)
			if err != nil {
				respondError(w, http.StatusNotFound, err)
				return
			}
			respondJSON(w, http.StatusOK, setting)
		case http.MethodPut:
			var req updateMerchantRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				respondError(w, http.StatusBadRequest, err)
				return
			}
			setting, err := merchantService.GetMerchantSetting(r.Context(), entityID)
			if err != nil {
				respondError(w, http.StatusNotFound, err)
				return
			}
			setting.CashboxID = req.CashboxID
			setting.Enabled = req.Enabled
			if err := merchantService.UpdateMerchantSetting(r.Context(), setting); err != nil {
				respondError(w, http.StatusInternalServerError, err)
				return
			}
			respondJSON(w, http.StatusOK, setting)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/api/v1/receipts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var req createReceiptRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}
		if req.OrderID == "" || req.CardID == "" || req.Amount <= 0 {
			respondError(w, http.StatusBadRequest, errors.New("order_id, card_id, and positive amount are required"))
			return
		}
		receipt, err := paymentService.CreateReceipt(r.Context(), req.OrderID, req.CardID, req.Amount)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}
		respondJSON(w, http.StatusCreated, receipt)
	})

	router.HandleFunc("/api/v1/receipts/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/v1/receipts/")
		if strings.HasSuffix(path, "/pay") {
			receiptID := strings.TrimSuffix(path, "/pay")
			receiptID = strings.TrimSuffix(receiptID, "/")
			if receiptID == "" {
				respondError(w, http.StatusBadRequest, errors.New("receipt id is required"))
				return
			}
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			if err := paymentService.PayReceipt(r.Context(), receiptID); err != nil {
				respondError(w, http.StatusBadRequest, err)
				return
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if strings.HasPrefix(path, "order/") {
			orderID := strings.TrimPrefix(path, "order/")
			if orderID == "" {
				respondError(w, http.StatusBadRequest, errors.New("order id is required"))
				return
			}
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			receipt, err := paymentService.GetReceiptByOrder(r.Context(), orderID)
			if err != nil {
				respondError(w, http.StatusNotFound, err)
				return
			}
			respondJSON(w, http.StatusOK, receipt)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})

	router.HandleFunc("/api/v1/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var req createTransactionRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				respondError(w, http.StatusBadRequest, err)
				return
			}
			if req.RestaurantID == "" || req.Currency == "" || req.Status == "" || req.CardID == "" || req.Amount <= 0 {
				respondError(w, http.StatusBadRequest, errors.New("restaurant_id, currency, status, card_id, and positive amount are required"))
				return
			}
			tx := &models.Transaction{
				ID:           req.ID,
				RestaurantID: req.RestaurantID,
				Currency:     req.Currency,
				Status:       req.Status,
				CreateTime:   req.CreateTime,
				PayTime:      req.PayTime,
				CancelTime:   req.CancelTime,
				CardID:       req.CardID,
				Amount:       req.Amount,
				Account:      req.Account,
			}
			if err := paymentService.SaveTransaction(r.Context(), tx); err != nil {
				respondError(w, http.StatusInternalServerError, err)
				return
			}
			respondJSON(w, http.StatusCreated, tx)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/api/v1/transactions/", func(w http.ResponseWriter, r *http.Request) {
		txID := strings.TrimPrefix(r.URL.Path, "/api/v1/transactions/")
		if txID == "" {
			respondError(w, http.StatusBadRequest, errors.New("transaction id is required"))
			return
		}
		switch r.Method {
		case http.MethodGet:
			tx, err := paymentService.GetTransaction(r.Context(), txID)
			if err != nil {
				respondError(w, http.StatusNotFound, err)
				return
			}
			respondJSON(w, http.StatusOK, tx)
		case http.MethodPut:
			var req createTransactionRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				respondError(w, http.StatusBadRequest, err)
				return
			}
			tx := &models.Transaction{
				ID:           txID,
				RestaurantID: req.RestaurantID,
				Currency:     req.Currency,
				Status:       req.Status,
				CreateTime:   req.CreateTime,
				PayTime:      req.PayTime,
				CancelTime:   req.CancelTime,
				CardID:       req.CardID,
				Amount:       req.Amount,
				Account:      req.Account,
			}
			if err := paymentService.UpdateTransaction(r.Context(), tx); err != nil {
				respondError(w, http.StatusInternalServerError, err)
				return
			}
			respondJSON(w, http.StatusOK, tx)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func respondJSON(w http.ResponseWriter, status int, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func respondError(w http.ResponseWriter, status int, err error) {
	respondJSON(w, status, errorResponse{Error: err.Error()})
}
