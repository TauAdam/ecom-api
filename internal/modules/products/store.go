package products

import "database/sql"

type Store struct {
	db *sql.DB
}

func NewProductsStore(db *sql.DB) Store {
	return Store{db: db}
}
