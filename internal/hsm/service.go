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
}

func (s *service) ARQCValidation(ctx context.Context, params ARQCParams) (bool, error) {
	hsmClient, err := createHsmClient(ctx)
	if err != nil {
		return false, err
	}

	return hsmClient.ARQCValidation(ctx, imk, params.PAN, params.PSN, params.ATC, params.ARQCMessage, params.ARQC, timeout)
}

func createHsmClient(ctx context.Context) (hsmLib.Client, error) {
	env := environment.NewLocal()
	lvl := zap.NewAtomicLevelAt(log.DebugLevel)
	logger := log.NewProductionLogger(&lvl)
	ctx = log.Context(ctx, logger)
	return hsmLib.New(ctx, env, "my_bari-hsm-lib-app")
}
