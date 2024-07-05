package service

import "github.com/egorus1442/Report-Generation-Microservice/pkg/repository"

type Autorization interface {
}

type SalesList interface {
}

type SalesItem interface {
}

type Service struct {
	Autorization
	SalesList
	SalesItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
