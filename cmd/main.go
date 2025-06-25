package main

import (
	"Billing-service-/internal/container"
)

func main() {
	app := container.Build()

	app.Run()
}
