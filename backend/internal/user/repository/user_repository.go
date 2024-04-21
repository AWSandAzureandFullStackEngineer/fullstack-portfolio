package repository

import (
	"backend/internal/user/model"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Save(user *model.User) error
	ExistsByEmail(email string) bool
}

type PostgresUserRepository struct {
	DB *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{DB: db}
}

func (repo *PostgresUserRepository) Save(user *model.User) error {
	query := `INSERT INTO users (id, first_name, last_name, username, email, password) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := repo.DB.Exec(query, user.ID, user.FirstName, user.LastName, user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("error saving user to database: %v", err)
	}
	return nil
}

func (repo *PostgresUserRepository) ExistsByEmail(email string) bool {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`
	var exists bool
	err := repo.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
