package service

import (
	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/egorus1442/Report-Generation-Microservice/pkg/repository"
)

type Autorization interface {
	CreateUser(user rgm.User) (int, error)
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
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
	}
}
