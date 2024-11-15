package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tommjj/tasks/internal/adapters/storage"
	"github.com/tommjj/tasks/internal/core/domain"
)

func newTestTaskRepository() (*taskRepository, error) {
	storage := storage.New("./data.json")

	return NewTaskRepository(storage)
}

func TestAddTask(t *testing.T) {
	repo, err := newTestTaskRepository()
	if err != nil {
		t.Fatal(err)
	}
	repo.AddTask(&domain.Task{
		Status:      domain.StatusDone,
		Title:       "test 01",
		Description: "title is my name",
	})
}

func TestGetTask(t *testing.T) {
	repo, err := newTestTaskRepository()
	if err != nil {
		t.Fatal(err)
	}
	data01, err := repo.GetTask(2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data01)

	data01.Title = "123"

	repo.AddTask(&domain.Task{
		Status:      domain.StatusDone,
		Title:       "test 01",
		Description: "title is my name",
	})

	data02, err := repo.GetTask(2)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data02)

	assert.NotEqual(t, data01.Title, data02.Title)
}

func TestGetTasks(t *testing.T) {
	repo, err := newTestTaskRepository()
	if err != nil {
		t.Fatal(err)
	}
	list01, err := repo.GetTasks()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", list01)

	if len(list01) == 0 {
		t.Skip("task slice is empty")
	}

	list01[0].Title = "123"

	list02, err := repo.GetTasks()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", list02)

	assert.NotEqual(t, list01[0].Title, list02[0].Title)
}

func TestUpdateTask(t *testing.T) {
	repo, err := newTestTaskRepository()
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.UpdateTask(&domain.Task{
		ID:    1,
		Title: "updated",
	})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, data.Title, "updated")
}

func TestDelTask(t *testing.T) {
	repo, err := newTestTaskRepository()
	if err != nil {
		t.Fatal(err)
	}

	list, err := repo.GetTasks()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", list)

	if len(list) == 0 {
		repo.AddTask(&domain.Task{
			Status:      domain.StatusDone,
			Title:       "test 01",
			Description: "title is my name",
		})
	}

	err = repo.DelTask(4)
	if err != nil {
		t.Fatal(err)
	}
}
