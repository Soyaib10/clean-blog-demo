package repository

import "github.com/Soyaib10/clean-blog-demo/internal/domain"


type CommentRepository interface {
    Create(comment *domain.Comment) error
    ListByPost(postID string) ([]*domain.Comment, error)
}
