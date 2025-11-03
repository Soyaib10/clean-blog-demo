package usecase

import (
	"time"

	"github.com/Soyaib10/clean-blog-demo/internal/domain"
	"github.com/Soyaib10/clean-blog-demo/internal/repository"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (uc *UserUsecase) CreateUser(name, email string) (*domain.User, error) {
	user := &domain.User{
		ID:        generateID(),
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

func (uc *UserUsecase) GetUser(id string) (*domain.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *UserUsecase) ListUsers() ([]*domain.User, error) {
	return uc.userRepo.List()
}
