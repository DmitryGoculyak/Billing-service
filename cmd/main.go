package main

import (
	"Billing-service-/internal/db"
	"Billing-service-/internal/service"
)

func main() {
	db.InitDB()
	service.RunServer()
}
