package repository

import "github.com/jmoiron/sqlx"

type Autorization interface {
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
	return &Repository{}
}
