package products

import (
	"github.com/TauAdam/ecom-api/internal/models"
	"github.com/TauAdam/ecom-api/shared/response"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	store models.ProductsStore
}

func NewHandler(store models.ProductsStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) InitRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetAllProducts).Methods(http.MethodGet)
}

func (h *Handler) handleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, err)
		return
	}
	err = response.SendJSON(w, http.StatusOK, products)
	if err != nil {
		log.Fatalf("failed to send response: %v", err)
	}
	return
}
