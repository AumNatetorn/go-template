package template

import (
	"database/sql"
	"fmt"
)

type Response struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Find(req Request) (*Response, error) {
	var resp Response
	row := r.db.QueryRow("SELECT name,age FROM profile WHERE id = ?", req.ID)
	err := row.Scan(&resp.Name, &resp.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("data not found:%v", err)
		}
		return nil, fmt.Errorf("error querying profile: %v", err)
	}

	return &resp, nil
}
