package repositories

import (
	"context"
	"database/sql"
	"go-api/api/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user entities.User) error
	FindAll(ctx context.Context) ([]entities.ResponseUser, error)
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
		"INSERT INTO users (name, email, username, password, address, phone, date_of_birth, gender, is_active) VALUES (?,?,?,?,?,?,?,?,?)",
		user.Name, user.Email, user.Username, user.Password, user.Address, user.Phone, user.DateOfBirth, user.Gender, user.IsActive,
	)
	return err
}

func (r *userRepository) FindAll(ctx context.Context) ([]entities.ResponseUser, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email, username, phone, gender, date_of_birth, is_active, address FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entities.ResponseUser
	for rows.Next() {
		var u entities.ResponseUser
		rows.Scan(&u.ID, 
			&u.Name, 
			&u.Email, 
			&u.Username,
			&u.Phone,
			&u.Gender,
			&u.DateOfBirth,
			&u.IsActive,
			&u.Address,
		)
		users = append(users, u)
	}

	return users, nil
}