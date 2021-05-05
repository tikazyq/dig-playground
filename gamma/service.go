package gamma

import (
	"dig-playground/interfaces"
	"log"
)

type Service struct {
	alphaSvc interfaces.AlphaService
	betaSvc  interfaces.BetaService
}

func (svc *Service) Run() {
	log.Println("gamma service run")
	svc.betaSvc.Setup()
	svc.alphaSvc.Init()
	svc.alphaSvc.Stop()
	svc.betaSvc.Cleanup()
}

func NewService(alphaSvc interfaces.AlphaService, betaSvc interfaces.BetaService) (interfaces.GammaService, error) {
	svc := &Service{}
	svc.alphaSvc = alphaSvc
	svc.betaSvc = betaSvc
	return svc, nil
}
