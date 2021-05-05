package beta

import (
	"dig-playground/inject"
	"dig-playground/interfaces"
	"go.uber.org/dig"
	"log"
)

func init() {
	_ = inject.Container.Provide(NewService, dig.Name("beta"))
}

type Service struct {
	alphaSvc interfaces.AlphaService
}

func (svc *Service) Setup() {
	log.Println("beta service setup")
}

func (svc *Service) Cleanup() {
	log.Println("beta service cleanup")
}

func NewService(alphaSvc interfaces.AlphaService) interfaces.BetaService {
	return &Service{alphaSvc: alphaSvc}
}
