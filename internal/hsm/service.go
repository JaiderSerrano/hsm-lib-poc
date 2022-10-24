package hsm

import (
	"context"

	"github.com/mercadolibre/fury_go-core/pkg/log"
	"github.com/mercadolibre/fury_hsm-lib/v3/environment"
	hsmLib "github.com/mercadolibre/fury_hsm-lib/v3/pkg/hsm"
	"go.uber.org/zap"
)

type service struct {
	Service
}

func NewHSMService() *service {
	return &service{}
}

type Service interface {
	// ARQCValidation validation of a ARQC cryptogram.
	ARQCValidation(ctx context.Context, params ARQCParams) (bool, error)
	// PINGeneration generate a new PIN and PVV.
	PINGeneration(ctx context.Context, pinGenerationParams PINGenerationParams) (map[string]string, error)
	// PVVGeneration generate a new PVV given a PIN.
	PVVGeneration(ctx context.Context, pinGenerationParams PVVGenerationParams) (map[string]string, error)
}

func (s *service) ARQCValidation(ctx context.Context, params ARQCParams) (bool, error) {
	hsmClient, err := createHsmClient(ctx)
	if err != nil {
		return false, err
	}

	return hsmClient.ARQCValidation(ctx, imk, params.PAN, params.PSN, params.ATC, params.ARQCMessage, params.ARQC, timeout)
}

func (s *service) PINGeneration(ctx context.Context, params PINGenerationParams) (map[string]string, error) {
	hsmClient, err := createHsmClient(ctx)
	if err != nil {
		return nil, err
	}

	return hsmClient.GeneratePIN(ctx, pek, pvkLeft, pvkRight, params.PAN, params.PVKI, timeout)
}

func (s *service) PVVGeneration(ctx context.Context, params PVVGenerationParams) (map[string]string, error) {
	hsmClient, err := createHsmClient(ctx)
	if err != nil {
		return nil, err
	}

	return hsmClient.GeneratePVV(ctx, pek, pvkLeft, pvkRight, params.PIN, params.PAN, params.PVKI, timeout)
}

func createHsmClient(ctx context.Context) (hsmLib.Client, error) {
	env := environment.NewLocal()
	lvl := zap.NewAtomicLevelAt(log.DebugLevel)
	logger := log.NewProductionLogger(&lvl)
	ctx = log.Context(ctx, logger)
	return hsmLib.New(ctx, env, "my_bari-hsm-lib-app")
}
