# Restaurant Payment Service

This service implements the card, merchant, and payment domains for the restaurant payment flow.

## Features

- Payment card storage and lifecycle
- Merchant settings management
- Receipt and transaction handling
- SQLite persistence via GORM
- Clean layered architecture with domain, application, infrastructure, and HTTP interface

## Run

1. `go mod tidy`
2. `go run main.go`

## HTTP API

- `POST /api/v1/cards`
- `GET /api/v1/cards/:id`
- `DELETE /api/v1/cards/:id`
- `POST /api/v1/merchants`
- `PUT /api/v1/merchants/:entityId`
- `GET /api/v1/merchants/:entityId`
- `POST /api/v1/receipts`
- `POST /api/v1/receipts/:id/pay`
- `GET /api/v1/receipts/order/:orderId`
- `POST /api/v1/transactions`
- `PUT /api/v1/transactions/:id`
- `GET /api/v1/transactions/:id`
