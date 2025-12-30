package services

import (
	"context"
	"go-api/api/entities"
	"testing"
)

type mockRepo struct{}

func (m mockRepo) Create(ctx context.Context, u entities.User) error {
	return nil
}
func (m mockRepo) FindAll(ctx context.Context) ([]entities.User, error) {
	return []entities.User{}, nil
}

func TestCreateUser(t *testing.T) {
	service := NewUserService(mockRepo{})
	err := service.Create(context.Background(), entities.User{
		Name: "Test",
	})
	if err != nil {
		t.Error("should not error")
	}
}