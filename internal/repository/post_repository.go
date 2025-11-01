package repository

import (
	"github.com/Soyaib10/clean-blog-demo/internal/domain"
)

type PostRepository interface {
	Create(post *domain.Comment) error
	GetByID(id string) (*domain.Post, error)
	ListByUser(userID string) ([]*domain.Post, error)
}
