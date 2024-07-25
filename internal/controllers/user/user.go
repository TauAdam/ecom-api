package user

import (
	"fmt"
	"github.com/TauAdam/ecom-api/internal/auth"
	"github.com/TauAdam/ecom-api/internal/controllers/response"
	"github.com/TauAdam/ecom-api/internal/models"
	"github.com/gorilla/mux"
	"log"
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

	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		response.SendError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already registered", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err)
	}
	err = h.store.CreateUser(models.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		response.SendError(w, http.StatusInternalServerError, err)
		return
	}

	err = response.SendJSON(w, http.StatusCreated, nil)
	if err != nil {
		log.Fatalf("could not send response: %v", err)
	}
}
