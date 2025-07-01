package service

import (
	repo "Billing-service-/internal/repository"
	proto "Billing-service-/pkg/proto"

	"context"
)

type BillingServiceServer interface {
	CreateWallets(ctx context.Context, req *proto.CreateWalletRequest) (*proto.WalletResponse, error)
	GetWallets(ctx context.Context, req *proto.GetWalletRequest) (*proto.WalletResponse, error)
}

type BillingServer struct {
	repo repo.BillingRepository
}

func BillingServerConstructor(
	repo repo.BillingRepository,
) *BillingServer {
	return &BillingServer{repo: repo}
}

func (s *BillingServer) CreateWallets(ctx context.Context, req *proto.CreateWalletRequest) (*proto.WalletResponse, error) {
	return s.repo.CreateWallet(ctx, req.UserId, req.CurrencyCode)
}

func (s *BillingServer) GetWallets(ctx context.Context, req *proto.GetWalletRequest) (*proto.WalletResponse, error) {
	return s.repo.GetWallet(ctx, req.UserId)
}
