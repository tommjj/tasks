package repository

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/tommjj/tasks/internal/core/domain"
	"github.com/tommjj/tasks/internal/core/ports"
)

type repo struct {
	file string
	m    sync.Mutex
}

func New(filename string) ports.IRepository {
	return &repo{
		file: filename,
	}
}

func (r *repo) Read() ([]domain.Task, error) {
	r.m.Lock()
	defer r.m.Unlock()

	data, err := os.ReadFile(r.file)
	if err != nil {
		return nil, err
	}

	tasks := []domain.Task{}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *repo) Sync(tasks []domain.Task) error {
	r.m.Lock()
	defer r.m.Unlock()

	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(r.file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
