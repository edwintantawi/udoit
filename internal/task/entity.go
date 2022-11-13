package task

import "time"

type ID string

type Task struct {
	ID          ID        `json:"id"`
	Content     string    `json:"content"`
	Description string    `json:"description"`
	IsDone      bool      `json:"id_done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
