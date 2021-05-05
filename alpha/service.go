package alpha

import (
	"dig-playground/interfaces"
	"log"
)

type Service struct {
}

func (svc *Service) Init() {
	log.Println("alpha service init")
}

func (svc *Service) Stop() {
	log.Println("alpha service stop")
}

func NewService() interfaces.AlphaService {
	return &Service{}
}
