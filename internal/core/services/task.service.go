package services

import (
	"github.com/tommjj/tasks/internal/core/domain"
	"github.com/tommjj/tasks/internal/core/ports"
)

type taskService struct {
	tasks        []domain.Task
	storage      ports.IStorage
	currentMaxID int
}

// func NewTaskService(repo ports.IStorage) (*taskService, error) {
// 	// tasks, err := repo.Read()
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// maxID := 0
// 	// if len(tasks) > 0 {
// 	// 	maxID = slices.MaxFunc(tasks, func(a, b domain.Task) int {
// 	// 		return cmp.Compare(a.ID, b.ID)
// 	// 	}).ID
// 	// }

// 	// return &taskService{
// 	// 	repo:         repo,
// 	// 	tasks:        tasks,
// 	// 	currentMaxID: maxID,
// 	// }, nil
// }

// func (t *taskService) AddTask(ctx context.Context, task *domain.Task) (*domain.Task, error) {

// }
