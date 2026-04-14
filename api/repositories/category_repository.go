package repositories

import (
	"context"
	"database/sql"
	"go-api/api/entities"
)

type CategoriesRepository interface {
	FindAllCategories(ctx context.Context, limit, offset int) ([]entities.Categories, int, error)
	FindByID(ctx context.Context, id int64) (*entities.Categories, error)
}

type categoriesRepository struct {
	db *sql.DB
}

func NewCategoriesRepository(db *sql.DB) CategoriesRepository {
	return &categoriesRepository{db}
}

func (r *categoriesRepository) FindAllCategories(
	ctx context.Context,
	limit, offset int,
) ([]entities.Categories, int, error){

	var total int
	err := r.db.QueryRowContext(
		ctx,
		"SELECT COUNT(*) FROM categories",
	).Scan(&total)
	if err != nil {
		return  nil, 0, err
	}

	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, name FROM categories LIMIT ? OFFSET ?",
		limit, offset,
	)
	if err != nil {
		return  nil, 0, err
	}
	defer rows.Close()

	categories := []entities.Categories{}
	for rows.Next(){
		var c entities.Categories
		if err := rows.Scan(
			&c.ID,
			&c.Name,
		); err != nil {
			return  nil, 0, err
		}
		categories = append(categories, c)
	}
	return categories, total, nil
}

func (r *categoriesRepository) FindByID(
	ctx context.Context,
	id int64,
) (*entities.Categories, error) {
	
	var c entities.Categories
	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, name FROM categories WHERE id = ?", id,
	).Scan(
		&c.ID,
		&c.Name,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &c, err
}