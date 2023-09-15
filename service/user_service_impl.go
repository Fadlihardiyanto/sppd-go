package service

import (
	"context"

	"github.com/Fadlihardiyanto/sppd-go/entity"
	"github.com/Fadlihardiyanto/sppd-go/models"
	"github.com/Fadlihardiyanto/sppd-go/repository"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type userService struct {
	DB             *gorm.DB
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(DB *gorm.DB, userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &userService{
		DB:             DB,
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (s *userService) Create(ctx context.Context, request models.CreateUserRequest) (models.CreateUserResponse, error) {
	var response models.CreateUserResponse

	// validate request
	err := s.Validate.Struct(request)
	if err != nil {
		return response, err
	}

	// insert to database
	userEntity := entity.UserEntity{
		IDUser:       request.IDUser,
		Username:     request.Username,
		Email:        request.Email,
		Password:     request.Password,
		Nama:         request.Nama,
		NomorTelepon: request.NomorTelepon,
		Alamat:       request.Alamat,
	}
	userEntity, err = s.UserRepository.Insert(ctx, userEntity)
	if err != nil {
		return response, err
	}

	// mapping response
	response = models.CreateUserResponse{
		Username:     userEntity.Username,
		Email:        userEntity.Email,
		Nama:         userEntity.Nama,
		NomorTelepon: userEntity.NomorTelepon,
		Alamat:       userEntity.Alamat,
	}

	return response, nil
}

func (s *userService) FindAll(ctx context.Context) ([]models.CreateUserResponse, error) {
	var responses []models.CreateUserResponse

	// get all data from database
	userEntities, err := s.UserRepository.FindAll(ctx)
	if err != nil {
		return responses, err
	}

	// mapping response
	for _, userEntity := range userEntities {
		response := models.CreateUserResponse{
			Username:     userEntity.Username,
			Email:        userEntity.Email,
			Nama:         userEntity.Nama,
			NomorTelepon: userEntity.NomorTelepon,
			Alamat:       userEntity.Alamat,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (s *userService) FindByID(ctx context.Context, id string) (models.CreateUserResponse, error) {
	var response models.CreateUserResponse

	// get data from database
	userEntity, err := s.UserRepository.FindByID(ctx, id)
	if err != nil {
		return response, err
	}
	// mapping response
	response = models.CreateUserResponse{
		Username:     userEntity.Username,
		Email:        userEntity.Email,
		Nama:         userEntity.Nama,
		NomorTelepon: userEntity.NomorTelepon,
		Alamat:       userEntity.Alamat,
	}

	return response, nil
}
