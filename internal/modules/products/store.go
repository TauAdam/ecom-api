package products

import (
	"database/sql"
	"github.com/TauAdam/ecom-api/internal/models"
)

type Store struct {
	db *sql.DB
}

func NewProductsStore(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) GetProducts() ([]models.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	products := make([]models.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoProducts(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}
	return products, nil
}

func scanRowIntoProducts(rows *sql.Rows) (*models.Product, error) {
	product := new(models.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Image,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}
