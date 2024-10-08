package cart

import (
	"database/sql"
	"github.com/TauAdam/ecom-api/internal/models"
)

type Store struct {
	db *sql.DB
}

func NewCartStore(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) CreateOrder(order models.Order) (int, error) {
	res, err := s.db.Exec("INSERT INTO orders (user_id, total, status, address) VALUES (?,?,?,?)", order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s Store) CreateOrderItem(orderItem models.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES (?,?,?,?)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	if err != nil {
		return err
	}
	return nil
}

func (s Store) UpdateProduct(product models.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = ?, description = ?, price = ?, image = ?, quantity = ? WHERE id = ?", product.Name, product.Description, product.Price, product.Image, product.Quantity, product.ID)
	if err != nil {
		return err
	}

	return nil
}
