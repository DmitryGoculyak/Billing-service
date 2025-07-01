package pgsql

import (
	"Billing-service-/internal/db/models"
	proto "Billing-service-/pkg/proto"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BillingRepo struct {
	db *sqlx.DB
}

func BillingRepositoryConstructor(
	db *sqlx.DB,
) *BillingRepo {
	return &BillingRepo{
		db: db,
	}
}

func (r *BillingRepo) CreateWallet(ctx context.Context, userID, currencyCode string) (*proto.WalletResponse, error) {
	var id string
	err := r.db.GetContext(ctx, &id, "INSERT INTO wallets (user_id,currency_code) VALUES ($1, $2) RETURNING id",
		userID, currencyCode)
	if err != nil {
		return nil, err
	}

	return &proto.WalletResponse{
		Id:           id,
		UserId:       userID,
		CurrencyCode: currencyCode,
		Balance:      0.0,
	}, nil
}

func (r *BillingRepo) GetWallet(ctx context.Context, userID string) (*proto.WalletResponse, error) {
	var wallet models.WalletDB
	err := r.db.GetContext(ctx, &wallet, "SELECT * FROM wallets WHERE user_id = $1", userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "wallet not found")
		}
		return nil, status.Error(codes.Internal, "unknown error")
	}
	return &proto.WalletResponse{
		Id:           wallet.Id,
		UserId:       wallet.UserId,
		CurrencyCode: wallet.CurrencyCode,
		Balance:      wallet.Balance,
	}, nil
}
