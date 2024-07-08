package repository

import (
	"fmt"

	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/jmoiron/sqlx"
)

type SalesListPostgres struct {
	db *sqlx.DB
}

func NewSalesListPostgres(db *sqlx.DB) *SalesListPostgres {
	return &SalesListPostgres{db: db}
}

func (r *SalesListPostgres) Create(userId int, list rgm.SalesList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, price, amount) VALUES ($1, $2, $3) RETURNING id", salesListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Price, list.Amount)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *SalesListPostgres) GetAll(userId int) ([]rgm.SalesList, error) {
	var lists []rgm.SalesList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.price, tl.amount FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		salesListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}
