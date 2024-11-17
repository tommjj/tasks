package repositories

import (
	"cmp"
	"context"
	"errors"
	"slices"
	"sync"
	"time"

	"github.com/tommjj/tasks/internal/core/domain"
	"github.com/tommjj/tasks/internal/core/ports"
)

type taskRepository struct {
	mutex        sync.Mutex
	tasks        []domain.Task
	currentMaxID int

	storage ports.IStorage
}

func NewTaskRepository(storage ports.IStorage) (ports.ITaskRepository, error) {
	tasks := []domain.Task{}

	err := storage.Load(context.Background(), &tasks)
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
func (t *taskRepository) AddTask(task *domain.Task) (*domain.Task, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.currentMaxID += 1

	newTask := domain.Task{
		ID:          t.currentMaxID,
		Status:      task.Status,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		CreatedAt:   time.Now(),
	}

	t.tasks = append(t.tasks, newTask)

	err := t.sync()
	if err != nil {
		return nil, err
	}

	return &newTask, nil
}

// GetTask
func (t *taskRepository) GetTask(id int) (*domain.Task, error) {
	index := slices.IndexFunc(t.tasks,
		func(v domain.Task) bool {
			return v.ID == id
		})

	if index == -1 {
		return nil, domain.ErrNotFound
	}

	task := t.tasks[index]
	return &task, nil
}

// GetTasks
func (t *taskRepository) GetTasks() ([]domain.Task, error) {
	tasks := slices.Clone(t.tasks)
	return tasks, nil
}

func (t *taskRepository) Swap(i, j int) error {
	leng := len(t.tasks)

	if i == j {
		return nil
	}

	if leng <= i || i < 0 {
		return nil
	}
	if leng <= j || j < 0 {
		return nil
	}

	t.tasks[i], t.tasks[j] = t.tasks[j], t.tasks[i]

	return t.sync()
}

// UpdateTask
func (t *taskRepository) UpdateTask(task *domain.Task) (*domain.Task, error) {
	index := slices.IndexFunc(t.tasks,
		func(v domain.Task) bool {
			return v.ID == task.ID
		})

	if index == -1 {
		return nil, domain.ErrNotFound
	}

	before := &t.tasks[index]

	if task.Title != "" {
		before.Title = task.Title
	}

	if task.Description != "" {
		before.Description = task.Description
	}

	if task.Priority != 0 {
		before.Priority = task.Priority
	}

	if task.Status != 0 {
		before.Status = task.Status
	}

	err := t.sync()
	if err != nil {
		return nil, err
	}

	after := *before
	return &after, nil
}

// DelTask
func (t *taskRepository) DelTask(id int) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	index := slices.IndexFunc(t.tasks,
		func(v domain.Task) bool {
			return v.ID == id
		})

	if index == -1 {
		return domain.ErrNotFound
	}

	t.tasks = slices.Delete(t.tasks, index, index+1)

	err := t.sync()
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) sync() error {
	err := t.storage.Sync(context.Background(), &t.tasks)
	return err
}
