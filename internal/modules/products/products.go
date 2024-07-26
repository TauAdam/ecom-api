package products

import (
	"errors"
	"fmt"
	"github.com/TauAdam/ecom-api/internal/models"
	"github.com/TauAdam/ecom-api/shared/request"
	"github.com/TauAdam/ecom-api/shared/response"
	"github.com/go-playground/validator/v10"
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
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleGetAllProducts(w http.ResponseWriter, _ *http.Request) {
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

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload models.Product
	if err := response.ParseJSON(r, &payload); err != nil {
		response.SendError(w, http.StatusBadRequest, err)
	}

	if err := request.Validate.Struct(payload); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		response.SendError(w, http.StatusBadRequest, fmt.Errorf("validation error: %v", validationErrors))
		return
	}

	err := h.store.CreateProduct(payload)
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, err)
		return
	}

	err = response.SendJSON(w, http.StatusOK, map[string]string{"result": "success"})
	if err != nil {
		log.Fatalf("failed to send response: %v", err)
	}
	return
}
