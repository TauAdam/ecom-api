package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/TauAdam/ecom-api/internal/models"
	"github.com/TauAdam/ecom-api/internal/modules/user"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockUserStore struct {
}

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := user.NewHandler(userStore)

	t.Run("should fail when user payload is invalid", func(t *testing.T) {
		payload := models.RegisterUserPayload{
			FirstName: "user",
			LastName:  "user",
			Email:     "some-invalid-email",
			Password:  "short",
		}
		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Errorf("error marshalling payload: %v", err)
		}
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})
	t.Run("should register user", func(t *testing.T) {
		payload := models.RegisterUserPayload{
			FirstName: "user",
			LastName:  "user",
			Email:     "valid@gmail.com",
			Password:  "valid-password",
		}
		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Errorf("error marshalling payload: %v", err)
		}
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.HandleRegister)
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			fmt.Print("message", rr.Body)
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}
	})
}
func (s *mockUserStore) CreateUser(payload models.User) error {
	return nil
}
func (s *mockUserStore) GetUserByID(id int) (*models.User, error) {
	return nil, nil
}
func (s *mockUserStore) GetUserByEmail(email string) (*models.User, error) {
	return nil, fmt.Errorf("not implemented")
}
