package server

import (
	cfg "Billing-service-/config"
	proto "Billing-service-/pkg/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func RunServer(cfg *cfg.GrpcServiceConfig, server proto.BillingServiceServer) {

	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterBillingServiceServer(grpcServer, server)

	log.Printf("[gRPC] Server started at time %v on port %v",
		time.Now().Format("[2006-01-02] [15:04]"), address)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
