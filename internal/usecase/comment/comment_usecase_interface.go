package comment

import "github.com/Soyaib10/clean-blog-demo/internal/domain"

type CommentUsecase interface {
	CreateComment(postID, userID, content string) (*domain.Comment, error)
	ListComments(postID string) ([]*domain.Comment, error)
}
