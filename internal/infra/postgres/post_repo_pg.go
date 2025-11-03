package postgres

import (
	"context"

	"github.com/Soyaib10/clean-blog-demo/internal/domain"
	"github.com/Soyaib10/clean-blog-demo/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postRepoPG struct {
    db *pgxpool.Pool
}

func NewPostRepo(db *pgxpool.Pool) repository.PostRepository {
    return &postRepoPG{db: db}
}

func (r *postRepoPG) Create(post *domain.Post) error {
    query := `INSERT INTO posts(id, user_id, title, content, created_at) VALUES($1,$2,$3,$4,$5)`
    _, err := r.db.Exec(context.Background(), query, post.ID, post.UserID, post.Title, post.Content, post.CreatedAt)
    return err
}

func (r *postRepoPG) GetByID(id string) (*domain.Post, error) {
    var post domain.Post
    query := `SELECT id, user_id, title, content, created_at FROM posts WHERE id=$1`
    err := r.db.QueryRow(context.Background(), query, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt)
    if err != nil {
        return nil, err
    }
    return &post, nil
}

func (r *postRepoPG) ListByUser(userID string) ([]*domain.Post, error) {
    rows, err := r.db.Query(context.Background(), `SELECT id, user_id, title, content, created_at FROM posts WHERE user_id=$1`, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts []*domain.Post
    for rows.Next() {
        var p domain.Post
        err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.CreatedAt)
        if err != nil {
            return nil, err
        }
        posts = append(posts, &p)
    }

    return posts, nil
}
