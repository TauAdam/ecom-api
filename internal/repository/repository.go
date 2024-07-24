package repository

import (
	"database/sql"
	"github.com/TauAdam/ecom-api/internal/models"
)

type Auth interface {
	GetUserByEmail(email string) (*models.User, error)
}
type Repository struct {
	Auth
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Auth: NewAuthMySQL(db),
	}
}
