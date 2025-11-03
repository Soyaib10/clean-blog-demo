package domain

import (
	"errors"
	"time"
)

type Comment struct {
	ID        string
	PostID    string
	UserID    string
	Content   string
	CreatedAt time.Time
}

// Validate ensures comment data is valid
func (c *Comment) Validate() error {
	if c.PostID == "" {
		return errors.New("post ID is required")
	}
	if c.UserID == "" {
		return errors.New("user ID is required")
	}
	if c.Content == "" {
		return errors.New("content is required")
	}
	return nil
}
