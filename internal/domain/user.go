package domain

import (
	"errors"
	"regexp"
)

type User struct {
	ID    string
	Name  string
	Email string
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}

	// simple email regex check
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !re.MatchString(u.Email) {
		return errors.New("invalid email format")
	}

	return nil
}
