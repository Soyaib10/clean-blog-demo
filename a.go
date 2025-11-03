package http

import (
	"github.com/Soyaib10/clean-blog-demo/internal/delivery/http/comment"
	"github.com/Soyaib10/clean-blog-demo/internal/delivery/http/post"
	"github.com/Soyaib10/clean-blog-demo/internal/delivery/http/user"
	"github.com/gorilla/mux"
)

type Handlers struct {
	User    *user.Handler
	Post    *post.Handler
	Comment *comment.Handler
}

func NewRouter(h *Handlers) *mux.Router {
	r := mux.NewRouter()

	userRouter := r.PathPrefix("/users").Subrouter()
	user.RegisterRoutes(userRouter, h.User)

	postRouter := r.PathPrefix("/posts").Subrouter()
	post.RegisterRoutes(postRouter, h.Post)

	commentRouter := r.PathPrefix("/comments").Subrouter()
	comment.RegisterRoutes(commentRouter, h.Comment)

	return r
}
