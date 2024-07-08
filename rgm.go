package rgm

type SalesList struct {
	Id     int     `json:"id" db:"id"`
	Title  string  `json:"title" db:"title" binding:"required"`
	Price  float64 `json:"price" db:"price"`
	Amount int     `json:"amount" db:"amount"`
}

type UsersList struct {
	Id     int
	UserId int
	ListID int
}
