package container

import (
	"Billing-service-/config"
	"Billing-service-/internal/db"
	"Billing-service-/internal/service"
	"go.uber.org/fx"
)

func Build() *fx.App {
	return fx.New(
		db.Module,
		config.Models,
		service.Module,
	)
}
