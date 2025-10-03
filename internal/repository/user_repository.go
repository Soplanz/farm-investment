package repository

import (
	"database/sql"
	"farm-investment/internal/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	Create(user models.User) (models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Username, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, user.Username, user.Email).Scan(&user.ID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
