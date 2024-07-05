package rgm

type SalesList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListID int
}

type SalesItem struct {
	Id       int     `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
	Done     bool    `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemID int
}
