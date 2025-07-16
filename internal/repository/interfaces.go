package repository

import (
	"Billing-service-/internal/entity"
	"context"
)

type BillingRepository interface {
	CreateWallet(ctx context.Context, userID, currencyCode string) (*entity.Wallet, error)
	GetWallet(ctx context.Context, userID string) (*entity.Wallet, error)
}
