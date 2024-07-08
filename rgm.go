package rgm

type SalesList struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	Amount int     `json:"amount"`
}

type UsersList struct {
	Id     int
	UserId int
	ListID int
}
