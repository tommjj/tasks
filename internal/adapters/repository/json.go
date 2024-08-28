package repository

import (
	"encoding/json"
	"errors"
	"io"
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

	var tasks []domain.Task

	file, err := os.Open(r.file)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(file)

	_, err = dec.Token()
	if err != nil {
		return nil, err
	}

	for dec.More() {
		var task domain.Task

		err := dec.Decode(&task)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return tasks, nil
			}
			return nil, err
		}
		tasks = append(tasks, task)
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
