package cart

import (
	"fmt"
	"github.com/TauAdam/ecom-api/internal/models"
)

func getCartItemsIDs(items []models.CartItem) ([]int, error) {
	productIDs := make([]int, len(items))
	for i, el := range items {
		if el.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product with id %d", el.ProductID)
		}
		productIDs[i] = el.ProductID
	}

	return productIDs, nil
}

func (h *Handler) createOrder(products []models.Product, cartItems []models.CartItem, userID int) (int, float64, error) {
	productsMap := make(map[int]models.Product)

	for _, p := range products {
		productsMap[p.ID] = p
	}

	if err := isItemsInStock(cartItems, productsMap); err != nil {
		return 0, 0, err
	}

	// TODO: calculate total price

	return 0, 0, nil
}

func isItemsInStock(items []models.CartItem, productsMap map[int]models.Product) error {
	if len(items) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range items {
		product, ok := productsMap[item.ProductID]

		if !ok {
			return fmt.Errorf("product with id %d is not available in the shop", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantitiy requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(cartItems []models.CartItem, products map[int]models.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}
	return total
}
