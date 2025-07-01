package service

import (
	"Billing-service-/internal/server"
	"go.uber.org/fx"
)

var Module = fx.Module("service",
	fx.Provide(
		BillingServerConstructor,
		func(s *BillingServer) BillingServiceServer { return s },
	),
	fx.Invoke(server.RunServer),
)
