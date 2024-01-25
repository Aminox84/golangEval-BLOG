package repositories

import (
	"blog/internal/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *models.User) (int64, error) {
	result, err := r.db.Exec(`
        INSERT INTO users (lastname, email) 
        VALUES (?, ?)`, user.Lastname, user.Email)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
