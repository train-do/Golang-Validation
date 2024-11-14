package repository

import (
	"database/sql"

	"github.com/train-do/Golang-Generics/model"
)

type RepoOrder struct {
	Db *sql.DB
}

func NewRepoOrder(db *sql.DB) *RepoOrder {
	return &RepoOrder{db}
}

func (r *RepoOrder) Insert(order *model.Order) error {
	query := `insert into "Order" ("destination_schedule_id", "name", "email", "phone", "message", "rating") values ($1, $2, $3, $4, $5, $6) returning id`
	err := r.Db.QueryRow(query, order.DestinationScheduleId, order.Name, order.Email, order.Phone, order.Message, order.Rating).Scan(&order.Id)
	if err != nil {
		return err
	}
	return nil
}
