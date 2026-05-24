package model

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (r CreateTaskRequest) Validate() error {
	if r.Title == "" {
		return fmt.Errorf("title is required")
	}
	return validateStatus(r.Status)
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (r UpdateTaskRequest) Validate() error {
	if r.Title == "" {
		return fmt.Errorf("title is required")
	}
	return validateStatus(r.Status)
}

func validateStatus(s string) error {
	switch s {
	case "", "todo", "in_progress", "done":
		return nil
	default:
		return fmt.Errorf("status must be one of: todo, in_progress, done")
	}
}
