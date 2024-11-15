package ports

import (
	"context"

	"github.com/tommjj/tasks/internal/core/domain"
)

type ITaskRepository interface {
	// Read is a func get all task form file
	Read() ([]domain.Task, error)
	// Sync is a func overwrite all task to file
	Sync([]domain.Task) error
}

type IStorage interface {
	Load(ctx context.Context, v any) error
	Sync(ctx context.Context, v any) error
}
