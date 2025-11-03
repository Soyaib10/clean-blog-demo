package repository

import "github.com/Soyaib10/clean-blog-demo/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	List() ([]*domain.User, error)
}