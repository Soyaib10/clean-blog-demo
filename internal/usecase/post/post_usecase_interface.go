package post

import (
	"github.com/Soyaib10/clean-blog-demo/internal/domain"
)

type PostUsecase interface {
	CreatePost(title, content, userID string) (*domain.Post, error)
	GetPost(id string) (*domain.Post, error)
	ListPostsByUser(userID string) ([]*domain.Post, error)
}
