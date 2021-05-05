package beta

import (
	"dig-playground/interfaces"
	"log"
)

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
