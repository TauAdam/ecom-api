package cart

import (
	"errors"
	"fmt"
	"github.com/TauAdam/ecom-api/internal/models"
	"github.com/TauAdam/ecom-api/shared/request"
	"github.com/TauAdam/ecom-api/shared/response"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	store         models.CartStore
	productsStore models.ProductsStore
}

func NewHandler(store models.CartStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) InitRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	var cart models.CartCheckoutPayload
	if err := response.ParseJSON(r, &cart); err != nil {
		response.SendError(w, http.StatusBadRequest, fmt.Errorf("invalid request payload %v", err))
		return
	}

	if err := request.Validate.Struct(cart); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		response.SendError(w, http.StatusBadRequest, fmt.Errorf("invalid request payload %v", validationErrors))
		return
	}

	productIDs, err := getCartItemsIDs(cart.Items)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, fmt.Errorf("invalid request payload %v", err))
		return
	}

	products, err := h.productsStore.GetProductByIDs(productIDs)
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, fmt.Errorf("failed to get products %v", err))
		return
	}

	orderID, totalPrice, err := h.createOrder(products, cart.Items, 1)
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, fmt.Errorf("failed to create order %v", err))
		return
	}

	err = response.SendJSON(w, http.StatusOK, map[string]any{
		"order_id": orderID,
		"total":    totalPrice,
	})
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, fmt.Errorf("failed to send response %v", err))
		return
	}

}
