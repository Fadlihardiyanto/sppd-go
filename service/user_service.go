package service

import (
	"context"

	"github.com/Fadlihardiyanto/sppd-go/models"
)

type UserService interface {
	Create(ctx context.Context, request models.CreateUserRequest) (models.CreateUserResponse, error)
	FindAll(ctx context.Context) ([]models.CreateUserResponse, error)
	FindByID(ctx context.Context, id string) (models.CreateUserResponse, error)
}
