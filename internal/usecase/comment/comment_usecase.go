package comment

import (
	"time"

	"github.com/Soyaib10/clean-blog-demo/internal/domain"
	"github.com/Soyaib10/clean-blog-demo/internal/repository"
	"github.com/Soyaib10/clean-blog-demo/internal/usecase/helpers"
)

type commentUsecase struct {
	commentRepo repository.CommentRepository
}

func NewCommentUsecase(commentRepo repository.CommentRepository) CommentUsecase {
	return &commentUsecase{commentRepo: commentRepo}
}

func (uc *commentUsecase) CreateComment(postID, userID, content string) (*domain.Comment, error) {
	comment := &domain.Comment{
		ID:        helpers.GenerateID(),
		PostID:    postID,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := comment.Validate(); err != nil {
		return nil, err
	}

	if err := uc.commentRepo.Create(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (uc *commentUsecase) ListComments(postID string) ([]*domain.Comment, error) {
	return uc.commentRepo.ListByPost(postID)
}
