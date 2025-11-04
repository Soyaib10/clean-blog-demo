package user

import "github.com/Soyaib10/clean-blog-demo/internal/domain"

type UserUsecase interface {
	CreateUser(name, email string) (*domain.User, error)
	GetUser(id string) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
}
