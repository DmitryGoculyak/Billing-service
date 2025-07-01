package handlers

import (
	"Billing-service-/internal/service"
	proto "Billing-service-/pkg/proto"
	"context"
)

type BillingHandler struct {
	proto.UnimplementedBillingServiceServer
	service service.BillingServiceServer
}

func BillingHandlerConstructor(
	service service.BillingServiceServer,
) *BillingHandler {
	return &BillingHandler{service: service}
}

func (h *BillingHandler) CreateWallet(ctx context.Context, req *proto.CreateWalletRequest) (*proto.WalletResponse, error) {
	return h.service.CreateWallets(ctx, req)
}

func (h *BillingHandler) GetWallet(ctx context.Context, req *proto.GetWalletRequest) (*proto.WalletResponse, error) {
	return h.service.GetWallets(ctx, req)
}
