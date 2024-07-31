package products

import (
	"database/sql"
	"fmt"
	"github.com/TauAdam/ecom-api/internal/models"
	"strings"
)

type Store struct {
	db *sql.DB
}

func (s Store) CreateProduct(payload models.Product) error {
	_, err := s.db.Exec("INSERT INTO products (Name, Description, Price, Image, Quantity) VALUES (?,?,?,?,?)", payload.Name, payload.Description, payload.Price, payload.Image, payload.Quantity)
	if err != nil {
		return err
	}

	return nil
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

func (s Store) GetProductByIDs(ids []int) ([]models.Product, error) {
	placeholders := strings.Repeat("?,", len(ids)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (%s?)", placeholders)

	args := make([]interface{}, len(ids))
	for i, v := range ids {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
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
