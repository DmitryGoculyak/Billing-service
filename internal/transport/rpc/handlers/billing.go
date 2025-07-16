package handlers

import (
	"Billing-service-/internal/service"
	proto "Billing-service-/pkg/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	user, err := h.service.CreateWallets(ctx, req.UserId, req.CurrencyCode)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create user: %v", err)
	}

	return &proto.WalletResponse{
		Id:           user.Id,
		UserId:       user.UserId,
		CurrencyCode: user.CurrencyCode,
		Balance:      0.00,
	}, nil
}

func (h *BillingHandler) GetWallet(ctx context.Context, req *proto.GetWalletRequest) (*proto.WalletResponse, error) {
	user, err := h.service.GetWallets(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get user: %v", err)
	}

	return &proto.WalletResponse{
		Id:           user.Id,
		UserId:       user.UserId,
		CurrencyCode: user.CurrencyCode,
		Balance:      user.Balance,
	}, nil
}
