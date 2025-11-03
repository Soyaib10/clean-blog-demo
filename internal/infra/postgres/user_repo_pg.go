package postgres

import (
	"context"

	"github.com/Soyaib10/clean-blog-demo/internal/domain"
	"github.com/Soyaib10/clean-blog-demo/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepoPG struct {
    db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) repository.UserRepository {
    return &userRepoPG{db: db}
}

func (r *userRepoPG) Create(user *domain.User) error {
    query := `INSERT INTO users(id, name, email, created_at) VALUES($1,$2,$3,$4)`
    _, err := r.db.Exec(context.Background(), query, user.ID, user.Name, user.Email, user.CreatedAt)
    return err
}

func (r *userRepoPG) GetByID(id string) (*domain.User, error) {
    var user domain.User
    query := `SELECT id, name, email, created_at FROM users WHERE id=$1`
    err := r.db.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepoPG) List() ([]*domain.User, error) {
    rows, err := r.db.Query(context.Background(), `SELECT id, name, email, created_at FROM users`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []*domain.User
    for rows.Next() {
        var u domain.User
        err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
        if err != nil {
            return nil, err
        }
        users = append(users, &u)
    }

    return users, nil
}
