package services

import (
	"context"
	"go-api/api/entities"
	"go-api/api/repositories"
)

type UserService interface {
	Create(ctx context.Context, user entities.User) error
	GetPaginated(ctx context.Context, page, limit int) ([]entities.ResponseUser, int, error)
	GetByID(ctx context.Context, id int64) (*entities.ResponseUser, error)
	Update(ctx context.Context, id int64, user entities.User) error
	Delete(ctx context.Context, id int64) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return  &userService{repo}
}

func (s *userService) Create(ctx context.Context, user entities.User) error {
	return s.repo.Create(ctx, user)
}

func (s *userService) GetPaginated(
	ctx context.Context,
	page, limit int,
) ([]entities.ResponseUser, int, error) {

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	return s.repo.FindWithPagination(ctx, limit, offset)
}

func (s *userService) GetByID(ctx context.Context, id int64) (*entities.ResponseUser, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *userService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *userService) Update(
	ctx context.Context,
	id int64,
	user entities.User,
) error {
	return s.repo.Update(ctx, id, user)
}
