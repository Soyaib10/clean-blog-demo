package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"

	router "github.com/Soyaib10/clean-blog-demo/internal/delivery/http"
	commentHandler "github.com/Soyaib10/clean-blog-demo/internal/delivery/http/comment"
	postHandler "github.com/Soyaib10/clean-blog-demo/internal/delivery/http/post"
	userHandler "github.com/Soyaib10/clean-blog-demo/internal/delivery/http/user"
	"github.com/Soyaib10/clean-blog-demo/internal/infra/postgres"
	"github.com/Soyaib10/clean-blog-demo/internal/usecase"
)

func main() {
	db, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/clean_blog_demo")
	if err != nil {
		log.Fatal(err)
	}

	// Repositories
	userRepo := postgres.NewUserRepo(db)
	postRepo := postgres.NewPostRepo(db)
	commentRepo := postgres.NewCommentRepo(db)

	// Usecases
	userUC := usecase.NewUserUsecase(userRepo)
	postUC := usecase.NewPostUsecase(postRepo)
	commentUC := usecase.NewCommentUsecase(commentRepo)

	// Handlers
	h := &router.Handlers{
		User:    userHandler.NewHandler(userUC),
		Post:    postHandler.NewHandler(postUC),
		Comment: commentHandler.NewHandler(commentUC),
	}

	r := router.NewRouter(h)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}