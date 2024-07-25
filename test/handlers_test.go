package test

import (
	"bytes"
	"encoding/json"
	"github.com/TauAdam/ecom-api/internal/controllers/user"
	"github.com/TauAdam/ecom-api/internal/models"
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
			Email:     "",
			Password:  "qwerty",
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
}
func (s *mockUserStore) CreateUser(payload models.User) error {
	return nil
}
func (s *mockUserStore) GetUserByID(id int) (*models.User, error) {
	return nil, nil
}
func (s *mockUserStore) GetUserByEmail(email string) (*models.User, error) {
	return nil, nil
}
