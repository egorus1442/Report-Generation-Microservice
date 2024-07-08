package service

import (
	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/egorus1442/Report-Generation-Microservice/pkg/repository"
)

type SalesListService struct {
	repo repository.SalesList
}

func NewSalesListService(repo repository.SalesList) *SalesListService {
	return &SalesListService{repo: repo}
}

func (s *SalesListService) Create(userId int, list rgm.SalesList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *SalesListService) GetAll(userId int) ([]rgm.SalesList, error) {
	return s.repo.GetAll(userId)
}
