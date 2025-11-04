package post

import (
	"time"

	"github.com/Soyaib10/clean-blog-demo/internal/domain"
	"github.com/Soyaib10/clean-blog-demo/internal/repository"
	"github.com/Soyaib10/clean-blog-demo/internal/usecase/helpers"
)

type postUsecase struct {
	postRepo repository.PostRepository
}

func NewPostUsecase(postRepo repository.PostRepository) PostUsecase {
	return &postUsecase{postRepo: postRepo}
}

func (uc *postUsecase) CreatePost(title, content, userID string) (*domain.Post, error) {
	post := &domain.Post{
		ID:        helpers.GenerateID(),
		Title:     title,
		Content:   content,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	if err := post.Validate(); err != nil {
		return nil, err
	}

	if err := uc.postRepo.Create(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (uc *postUsecase) GetPost(id string) (*domain.Post, error) {
	return uc.postRepo.GetByID(id)
}

func (uc *postUsecase) ListPostsByUser(userID string) ([]*domain.Post, error) {
	return uc.postRepo.ListByUser(userID)
}
