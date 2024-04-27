package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var (
	ErrTitleIsRequired       = errors.New("title is required")
	ErrDescriptionIsRequired = errors.New("description is required")
	ErrCategoryIsRequired    = errors.New("category is required")
)

func NewCourse(title, description, category string) (*Course, error) {
	course := &Course{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Category:    category,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := course.Validate()
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (p *Course) Validate() error {
	if p.Title == "" {
		return ErrTitleIsRequired
	}

	if p.Description == "" {
		return ErrDescriptionIsRequired
	}

	if p.Category == "" {
		return ErrCategoryIsRequired
	}

	return nil
}
