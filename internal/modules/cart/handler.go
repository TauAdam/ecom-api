package cart

import (
	"github.com/TauAdam/ecom-api/internal/models"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	store models.CartStore
}

func NewHandler(store models.CartStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) InitRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods(http.MethodPost)
}
