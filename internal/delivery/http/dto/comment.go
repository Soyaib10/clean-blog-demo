package dto

type CreateCommentRequest struct {
	PostID  string `json:"post_id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

type CommentResponse struct {
	ID        string `json:"id"`
	PostID    string `json:"post_id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
