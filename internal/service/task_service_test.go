package service

import (
	"context"
	"testing"

	"github.com/edvina16/icpal-backend/internal/database"
)

type mockTaskDb struct{}

func (m *mockTaskDb) ListTasks(ctx context.Context) ([]database.Task, error) {
	return []database.Task{
		{ID: 1, Title: "Test", Content: "Content test"},
	}, nil
}

func (m *mockTaskDb) CreateTask(ctx context.Context, params database.CreateTaskParams) (database.Task, error) {
	return database.Task{ID: 2, Title: params.Title, Content: params.Content}, nil
}

func TestListTasks(t *testing.T) {
	svc := TaskService{DB: &mockTaskDb{}}
	tasks, err := svc.ListTasks(context.Background())
	if err != nil {
		t.Fatalf("unexpected error listing tasks: %v", err)
	}
	if len(tasks) != 1 || tasks[0].Title != "Test" {
		t.Fatalf("expected one task with the title 'Test', got %v", tasks)
	}
}

func TestCreateTask(t *testing.T) {
	svc := TaskService{DB: &mockTaskDb{}}
	task, err := svc.CreateTask(context.Background(), "New", "Content")
	if err != nil {
		t.Fatalf("unexpected error creating task: %v", err)
	}
	if task.Title != "New" || task.Content != "Content" {
		t.Fatalf("unexpected task: %+v", task)
	}
}
