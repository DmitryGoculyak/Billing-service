package service

import (
	"Billing-service-/internal/entity"
	repo "Billing-service-/internal/repository"
	"context"
)

type BillingServiceServer interface {
	CreateWallets(ctx context.Context, userID, currencyName string) (*entity.Wallet, error)
	GetWallets(ctx context.Context, userID string) (*entity.Wallet, error)
}

type BillingServer struct {
	repo repo.BillingRepository
}

func BillingServerConstructor(
	repo repo.BillingRepository,
) *BillingServer {
	return &BillingServer{repo: repo}
}

func (s *BillingServer) CreateWallets(ctx context.Context, UserID, currencyName string) (*entity.Wallet, error) {
	return s.repo.CreateWallet(ctx, UserID, currencyName)
}

func (s *BillingServer) GetWallets(ctx context.Context, userID string) (*entity.Wallet, error) {
	return s.repo.GetWallet(ctx, userID)
}
