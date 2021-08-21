package service1

import (
	"goboss/internal/app/app1/service"
)

type Service struct {
	*service.Service
}

func NewService(s *service.Service) *Service {
	return &Service{
		Service: s,
	}
}
