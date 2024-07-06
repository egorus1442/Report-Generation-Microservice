package repository

import (
	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user rgm.User) (int, error)
}

type SalesList interface {
}

type SalesItem interface {
}

type Repository struct {
	Autorization
	SalesList
	SalesItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
	}
}
