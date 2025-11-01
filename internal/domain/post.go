package domain

import "errors"

type Post struct {
	ID       string
	Title    string
	Content  string
	AuthorID string
}

// Validate ensures post data is valid
func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}
	if p.Content == "" {
		return errors.New("content is required")
	}
	if p.AuthorID == "" {
		return errors.New("author ID is required")
	}
	return nil
}
