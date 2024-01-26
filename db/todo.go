package db

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UserId      string    `json:"userId"`
}

func NewTodo(description string) *Todo {
	return &Todo{
		ID:          uuid.New(),
		Completed:   false,
		Description: description,
		CreatedAt:   time.Now(),
	}
}

func (t *Todo) Update(completed bool, description string) {
	t.Completed = completed
	t.Description = description
}
