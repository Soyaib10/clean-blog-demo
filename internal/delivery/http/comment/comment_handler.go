package comment

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Soyaib10/clean-blog-demo/internal/delivery/http/dto"
	"github.com/Soyaib10/clean-blog-demo/internal/usecase"
	"github.com/gorilla/mux"
)

type Handler struct {
	uc usecase.CommentUsecase
}

func NewHandler(uc usecase.CommentUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

// CreateComment handler
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	comment, err := h.uc.CreateComment(req.PostID, req.UserID, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.CommentResponse{
		ID:        comment.ID,
		PostID:    comment.PostID,
		UserID:    comment.UserID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.Format(time.RFC3339),
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GetComment handler
func (h *Handler) ListByPost(w http.ResponseWriter, r *http.Request) {
	postID := mux.Vars(r)["postID"]
	comments, err := h.uc.ListComments(postID)
	if err != nil {
		http.Error(w, "could not fetch comments", http.StatusInternalServerError)
		return
	}

	var resp []dto.CommentResponse
	for _, c := range comments {
		resp = append(resp, dto.CommentResponse{
			ID:        c.ID,
			PostID:    c.PostID,
			UserID:    c.UserID,
			Content:   c.Content,
			CreatedAt: c.CreatedAt.Format(time.RFC3339),
		})
	}

	json.NewEncoder(w).Encode(resp)
}
