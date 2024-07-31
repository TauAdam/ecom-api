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
