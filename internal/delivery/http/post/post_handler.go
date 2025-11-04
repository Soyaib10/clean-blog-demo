package post

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Soyaib10/clean-blog-demo/internal/delivery/http/dto"
	"github.com/Soyaib10/clean-blog-demo/internal/usecase/post"
	"github.com/gorilla/mux"
)

type Handler struct {
	uc post.PostUsecase
}

func NewHandler(uc post.PostUsecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

// CreatePost handler
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	post, err := h.uc.CreatePost(req.UserID, req.Title, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.PostResponse{
		ID:        post.ID,
		UserID:    post.UserID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// GetPost handler
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	post, err := h.uc.GetPost(id)
	if err != nil {
		http.Error(w, "post not found", http.StatusNotFound)
		return
	}

	    resp := dto.PostResponse{
	        ID:        post.ID,
	        UserID:    post.UserID,
	        Title:     post.Title,
	        Content:   post.Content,
	        CreatedAt: post.CreatedAt.Format(time.RFC3339),
	    }
	
	    json.NewEncoder(w).Encode(resp)
	}
	
	// ListByUser handler
	func (h *Handler) ListByUser(w http.ResponseWriter, r *http.Request) {
		userID := mux.Vars(r)["userID"]
		posts, err := h.uc.ListPostsByUser(userID)
		if err != nil {
			http.Error(w, "could not fetch posts", http.StatusInternalServerError)
			return
		}
	
		var resp []dto.PostResponse
		for _, p := range posts {
			resp = append(resp, dto.PostResponse{
				ID:        p.ID,
				UserID:    p.UserID,
				Title:     p.Title,
				Content:   p.Content,
				CreatedAt: p.CreatedAt.Format(time.RFC3339),
			})
		}
	
		json.NewEncoder(w).Encode(resp)
	}
