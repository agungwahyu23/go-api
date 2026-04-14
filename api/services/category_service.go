package services

import (
	"context"
	"go-api/api/entities"
	"go-api/api/repositories"
)

type CategoriesService interface {
	GetAllCategories(ctx context.Context, page, limit int) ([]entities.Categories, int, error)
	GetByIDCategories(ctx context.Context, id int64) (*entities.Categories, error)
}

type categoriesService struct {
	repo repositories.CategoriesRepository
}

func NewCategoriesService(repo repositories.CategoriesRepository) CategoriesService {
	return &categoriesService{repo}
}

func (s *categoriesService) GetAllCategories(
	ctx context.Context,
	page, limit int,
) ([]entities.Categories, int, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	return s.repo.FindAllCategories(ctx, limit, offset)
}

func (s *categoriesService) GetByIDCategories(ctx context.Context, id int64) (*entities.Categories, error) {
	return s.repo.FindByID(ctx, id)
}