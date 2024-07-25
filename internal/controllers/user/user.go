package user

import (
	"github.com/TauAdam/ecom-api/internal/controllers/response"
	"github.com/TauAdam/ecom-api/internal/models"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	store models.UserStore
}

func NewHandler(store models.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) InitRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload models.RegisterUserPayload
	if err := response.ParseJSON(r, &payload); err != nil {
		response.SendError(w, http.StatusBadRequest, err)
	}

}
