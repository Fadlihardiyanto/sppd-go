package repository

import (
	"context"

	"github.com/Fadlihardiyanto/sppd-go/entity"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (r *userRepositoryImpl) Insert(ctx context.Context, userEntity entity.UserEntity) (entity.UserEntity, error) {

	result := r.db.Create(&userEntity)
	if result.Error != nil {
		return userEntity, result.Error
	}

	return userEntity, nil
}

func (r *userRepositoryImpl) FindAll(ctx context.Context) ([]entity.UserEntity, error) {
	var users []entity.UserEntity
	result := r.db.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id string) (entity.UserEntity, error) {
	var user entity.UserEntity
	SQL := "SELECT * FROM users WHERE id_user = ?"
	result := r.db.Raw(SQL, id).Scan(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (r *userRepositoryImpl) Update(ctx context.Context, userEntity entity.UserEntity) (entity.UserEntity, error) {
	SQL := "UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?"
	result := r.db.Exec(SQL, userEntity.Username, userEntity.Email, userEntity.Password, userEntity.IDUser)
	if result.Error != nil {
		return userEntity, result.Error
	}

	return userEntity, nil
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id int) (entity.UserEntity, error) {
	var user entity.UserEntity
	SQL := "DELETE FROM users WHERE id = ?"
	result := r.db.Raw(SQL, id).Scan(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
