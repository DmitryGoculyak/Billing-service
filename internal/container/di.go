package container

import (
	"Billing-service-/config"
	"Billing-service-/internal/db"
	repo "Billing-service-/internal/repository/pgsql"
	"Billing-service-/internal/service"
	"Billing-service-/internal/transport/rpc/handlers"
	"go.uber.org/fx"
)

func Build() *fx.App {
	return fx.New(
		db.Module,
		config.Models,
		service.Module,
		handlers.Module,
		repo.Module,
	)
}
