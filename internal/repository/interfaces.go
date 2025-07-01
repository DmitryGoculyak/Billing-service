package repository

import (
	proto "Billing-service-/pkg/proto"
	"context"
)

type BillingRepository interface {
	CreateWallet(ctx context.Context, userID, currencyCode string) (*proto.WalletResponse, error)
	GetWallet(ctx context.Context, userID string) (*proto.WalletResponse, error)
}
