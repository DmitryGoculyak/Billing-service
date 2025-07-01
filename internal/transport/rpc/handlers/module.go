package handlers

import (
	proto "Billing-service-/pkg/proto"
	"go.uber.org/fx"
)

var Module = fx.Module("handlers",
	fx.Provide(
		BillingHandlerConstructor,
		func(h *BillingHandler) proto.BillingServiceServer { return h },
	),
)
