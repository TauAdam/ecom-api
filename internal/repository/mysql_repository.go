package repository

import (
	"database/sql"
	"github.com/TauAdam/ecom-api/internal/models"
)

type AuthMySQL struct {
	db *sql.DB
}

func NewAuthMySQL(db *sql.DB) *AuthMySQL {
	return &AuthMySQL{db: db}
}

func (s *AuthMySQL) GetUserByEmail(email string) (*models.User, error) {

}
