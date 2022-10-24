package main

import (
	"context"

	"github.com/mercadolibre/fury_go-core/pkg/log"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
	"github.com/mercadolibre/hsm-lib-poc/internal/hsm"
	"github.com/mercadolibre/hsm-lib-poc/internal/hsm/handler"
)

func main() {
	if err := run(); err != nil {
		log.Error(context.Background(), "cannot run application", log.Err(err))
	}
}

func run() error {
	// Start fury application
	app, err := fury.NewWebApplication()
	if err != nil {
		return err
	}

	// Handler creation
	hsmService := hsm.NewHSMService()
	hsmHandler := handler.NewHSMHandler(hsmService)

	// HSM functionalities
	app.Post("/hsm", hsmHandler.ARQCValidation)

	return app.Run()
}