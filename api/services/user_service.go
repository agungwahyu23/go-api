package services

import (
	"context"
	"go-api/api/entities"
	"go-api/api/repositories"
)

type UserService interface {
	Create(ctx context.Context, user entities.User) error
	GetAll(ctx context.Context) ([]entities.User, error)
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

func (s *userService) GetAll(ctx context.Context) ([]entities.User, error) {
	return s.repo.FindAll(ctx)
}