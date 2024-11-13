package domain

import "time"

type Task struct {
	ID          int       `json:"id"`
	Status      Status    `json:"status,omitempty"`
	Priority    Priority  `json:"priority,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
