package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
