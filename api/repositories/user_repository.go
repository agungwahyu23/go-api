package repositories

import (
	"context"
	"database/sql"
	"go-api/api/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user entities.User) error
	FindAll(ctx context.Context) ([]entities.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user entities.User) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO users (name, email) VALUES (?, ?)",
		user.Name, user.Email,
	)
	return err
}

func (r *userRepository) FindAll(ctx context.Context) ([]entities.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var u entities.User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}

	return users, nil
}