package usecase

import (
	"time"

	"github.com/Soyaib10/clean-blog-demo/internal/domain"
	"github.com/Soyaib10/clean-blog-demo/internal/repository"
)

type PostUsecase struct {
	postRepo repository.PostRepository
}

func NewPostUsecase(postRepo repository.PostRepository) *PostUsecase {
	return &PostUsecase{postRepo: postRepo}
}

func (uc *PostUsecase) CreatePost(title, content, userID string) (*domain.Post, error) {
	post := &domain.Post{
		ID:        generateID(),
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

func (uc *PostUsecase) GetPost(id string) (*domain.Post, error) {
	return uc.postRepo.GetByID(id)
}

func (uc *PostUsecase) ListPostsByUser(userID string) ([]*domain.Post, error) {
	return uc.postRepo.ListByUser(userID)
}
