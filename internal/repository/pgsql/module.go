package pgsql

import (
	repo "Billing-service-/internal/repository"
	"go.uber.org/fx"
)

var Module = fx.Module("pgsql",
	fx.Provide(
		BillingRepositoryConstructor,
		func(r *BillingRepo) repo.BillingRepository { return r },
	),
)
