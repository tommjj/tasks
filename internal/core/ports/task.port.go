package ports

import "github.com/tommjj/tasks/internal/core/domain"

type IRepository interface {
	// Read is a func get all task form file
	Read() ([]domain.Task, error)
	// Sync is a func overwrite all task to file
	Sync([]domain.Task) error
}
