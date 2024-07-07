package service

import (
	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/egorus1442/Report-Generation-Microservice/pkg/repository"
)

type Authorization interface {
	CreateUser(user rgm.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type SalesList interface {
}

type SalesItem interface {
}

type Service struct {
	Authorization
	SalesList
	SalesItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
