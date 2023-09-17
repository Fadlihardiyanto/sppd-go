package repository

import (
	"context"

	"github.com/Fadlihardiyanto/sppd-go/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, userEntity entity.UserEntity) (entity.UserEntity, error)
	FindAll(ctx context.Context) ([]entity.UserEntity, error)
	FindByID(ctx context.Context, id string) (entity.UserEntity, error)
	Update(ctx context.Context, userEntity entity.UserEntity) (entity.UserEntity, error)
	Delete(ctx context.Context, id string) (entity.UserEntity, error)
	Login(ctx context.Context, username string) (entity.UserEntity, error)
}
