package ports

import (
	"context"

	"github.com/tommjj/tasks/internal/core/domain"
)

type ITaskRepository interface {
	AddTask(task *domain.Task) (*domain.Task, error)
	GetTask(id int) (*domain.Task, error)
	GetTasks() ([]domain.Task, error)
	Swap(i, j int) error
	UpdateTask(task *domain.Task) (*domain.Task, error)
	DelTask(id int) error
}

type IStorage interface {
	Load(ctx context.Context, v any) error
	Sync(ctx context.Context, v any) error
}
