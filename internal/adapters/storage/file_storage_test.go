package storage

import (
	"testing"
	"time"

	"github.com/tommjj/tasks/internal/core/domain"
)

func TestWrite(t *testing.T) {
	r := New("./data.json")

	err := r.Sync([]domain.Task{
		{
			ID:          1,
			Status:      domain.StatusDone,
			Title:       "test 01",
			Description: "title is my name",
			CreatedAt:   time.Now(),
		},
		{
			ID:          2,
			Priority:    domain.Pri4,
			Title:       "test 02",
			Description: "title is my name",
			CreatedAt:   time.Now(),
		},
	})
	if err != nil {
		t.Error(err)
	}
}

func TestRead(t *testing.T) {
	r := New("./data.json")

	tasks := []domain.Task{}
	err := r.Load(&tasks)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", tasks)
}
