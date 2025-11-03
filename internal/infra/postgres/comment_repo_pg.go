package postgres

import (
	"context"

	"github.com/Soyaib10/clean-blog-demo/internal/domain"
	"github.com/Soyaib10/clean-blog-demo/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type commentRepoPG struct {
    db *pgxpool.Pool
}

func NewCommentRepo(db *pgxpool.Pool) repository.CommentRepository {
    return &commentRepoPG{db: db}
}

func (r *commentRepoPG) Create(comment *domain.Comment) error {
    query := `INSERT INTO comments(id, post_id, user_id, content, created_at) VALUES($1,$2,$3,$4,$5)`
    _, err := r.db.Exec(context.Background(), query, comment.ID, comment.PostID, comment.UserID, comment.Content, comment.CreatedAt)
    return err
}

func (r *commentRepoPG) ListByPost(postID string) ([]*domain.Comment, error) {
    rows, err := r.db.Query(context.Background(), `SELECT id, post_id, user_id, content, created_at FROM comments WHERE post_id=$1`, postID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var comments []*domain.Comment
    for rows.Next() {
        var c domain.Comment
        err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt)
        if err != nil {
            return nil, err
        }
        comments = append(comments, &c)
    }

    return comments, nil
}
