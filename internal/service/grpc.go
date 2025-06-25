package service

import (
	"Billing-service-/internal/db"
	"Billing-service-/internal/db/models"
	proto "Billing-service-/pkg/proto"
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

type BillingServer struct {
	proto.UnimplementedBillingServiceServer
}

func (s *BillingServer) CreateWallet(ctx context.Context, req *proto.CreateWalletRequest) (*proto.WalletResponse, error) {
	var id string
	err := db.DB.Get(&id, "INSERT INTO wallets (user_id,currency_code) VALUES ($1, $2) RETURNING id;",
		req.UserId, req.CurrencyCode)
	if err != nil {
		return nil, err
	}
	return &proto.WalletResponse{
		Id:           id,
		UserId:       req.UserId,
		CurrencyCode: req.CurrencyCode,
		Balance:      0.0,
	}, nil
}

func (s *BillingServer) GetWallet(ctx context.Context, req *proto.GetWalletRequest) (*proto.WalletResponse, error) {
	var wallet models.WalletDB
	err := db.DB.Get(&wallet, "SELECT * FROM wallets WHERE user_id = $1", req.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("Error:", err)
			return nil, status.Errorf(codes.NotFound, "wallet not found")
		}
		log.Println("Error:", err)
		return nil, status.Error(codes.Internal, "unknown error")
	}
	return &proto.WalletResponse{
		Id:           wallet.Id,
		UserId:       wallet.UserId,
		CurrencyCode: wallet.CurrencyCode,
		Balance:      wallet.Balance,
	}, nil
}

func (s *BillingServer) DeleteAllWallets(ctx context.Context, _ *proto.Empty) (*proto.DeleteWallets, error) {
	_, err := db.DB.Exec("DELETE FROM wallets")
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}
	return &proto.DeleteWallets{Message: "All wallets deleted"}, nil
}

func RunServer() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterBillingServiceServer(grpcServer, &BillingServer{})

	log.Printf("[gRPC] Server started at time %v on port %v", time.Now().Format("[2006-01-02] [15:04]"), lis.Addr().(*net.TCPAddr).Port)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
