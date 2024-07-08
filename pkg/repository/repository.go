package repository

import (
	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user rgm.User) (int, error)
	GetUser(username, password string) (rgm.User, error)
}

type SalesList interface {
}

type Repository struct {
	Authorization
	SalesList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
