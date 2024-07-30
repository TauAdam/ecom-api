package user

import (
	"database/sql"
	"fmt"
	"github.com/TauAdam/ecom-api/internal/models"
)

type Store struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := s.db.QueryRow(
		"SELECT * FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func scanRowIntoUser(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)

	err := rows.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.FirstName,
		&user.LastName,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s Store) CreateUser(payload models.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VAlUES (?,?,?,?)", payload.FirstName, payload.LastName, payload.Email, payload.Password)
	if err != nil {
		return err
	}

	return nil
}
func (s Store) GetUserByID(id int) (*models.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	user := new(models.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}
