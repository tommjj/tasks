package repositories

import (
	"cmp"
	"errors"
	"slices"

	"github.com/tommjj/tasks/internal/core/domain"
	"github.com/tommjj/tasks/internal/core/ports"
)

type taskRepository struct {
	tasks        []domain.Task
	currentMaxID int
	storage      ports.IStorage
}

func NewTaskRepository(storage ports.IStorage) (*taskRepository, error) {
	tasks := []domain.Task{}

	err := storage.Load(&tasks)
	if err != nil && !errors.Is(err, domain.ErrNotExist) {
		return nil, err
	}

	currentMaxID := 0

	if len(tasks) > 0 {
		currentMaxID = slices.MaxFunc(tasks,
			func(a, b domain.Task) int {
				return cmp.Compare(a.ID, b.ID)
			}).ID
	}

	return &taskRepository{
		tasks:        tasks,
		currentMaxID: currentMaxID,
		storage:      storage,
	}, nil
}

// AddTask
func (t *taskRepository) AddTask() {

}

// GetTask

// GetTasks

// UpdateTask

// DelTask
