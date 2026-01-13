package repositories

import (
	"context"
	"database/sql"
	"go-api/api/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user entities.User) error
	FindWithPagination(ctx context.Context, limit, offset int) ([]entities.ResponseUser, int, error)
	FindByID(ctx context.Context, id int64) (*entities.ResponseUser, error)
	Update(ctx context.Context, id int64, user entities.User) error
	Delete(ctx context.Context, id int64) error
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

func (r *userRepository) FindWithPagination(
	ctx context.Context,
	limit, offset int,
) ([]entities.ResponseUser, int, error) {

	var total int
	err := r.db.QueryRowContext(
		ctx,
		"SELECT COUNT(*) FROM users",
	).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name, email, username, phone, gender, date_of_birth, is_active, address FROM users LIMIT ? OFFSET ?",
		limit, offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	users := []entities.ResponseUser{}
	for rows.Next() {
		var u entities.ResponseUser
		if err := rows.Scan(
			&u.ID, 
			&u.Name, 
			&u.Email, 
			&u.Username,
			&u.Phone,
			&u.Gender,
			&u.DateOfBirth,
			&u.IsActive,
			&u.Address); err != nil {
			return nil, 0, err
		}
		users = append(users, u)
	}

	return users, total, nil
}

func (r *userRepository) FindByID(
	ctx context.Context,
	id int64,
) (*entities.ResponseUser, error) {
	
	var u entities.ResponseUser
	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, name, email, username, phone, gender, date_of_birth, is_active, address FROM users WHERE id = ?", id,
	).Scan(
		&u.ID, 
		&u.Name, 
		&u.Email, 
		&u.Username,
		&u.Phone,
		&u.Gender,
		&u.DateOfBirth,
		&u.IsActive,
		&u.Address,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &u, err
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(
		ctx,
		"DELETE FROM users WHERE id = ?",
		id,
	)
	return err
}

func (r *userRepository) Update(
	ctx context.Context,
	id int64,
	user entities.User,
) error {
	_, err := r.db.ExecContext(
		ctx,
		"UPDATE users SET name = ?, email = ?, username = ?, address = ?, phone = ?, date_of_birth = ?, gender = ?, is_active = ? WHERE id = ?",
		user.Name,
		user.Email,
		user.Username,
		user.Address,
		user.Phone,
		user.DateOfBirth,
		user.Gender,
		user.IsActive,
		id,
	)
	return err
}
