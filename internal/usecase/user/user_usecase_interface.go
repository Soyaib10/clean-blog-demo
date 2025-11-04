package user

import (
	"time"

	"github.com/Soyaib10/clean-blog-demo/internal/domain"
	"github.com/Soyaib10/clean-blog-demo/internal/repository"
	"github.com/Soyaib10/clean-blog-demo/internal/usecase/helpers"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (uc *userUsecase) CreateUser(name, email string) (*domain.User, error) {
	user := &domain.User{
		ID:        helpers.GenerateID(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUsecase) GetUser(id string) (*domain.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *userUsecase) ListUsers() ([]*domain.User, error) {
	return uc.userRepo.List()
}