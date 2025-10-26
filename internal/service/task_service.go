package service

import (
	"context"

	"github.com/edvina16/icpal-backend/internal/database"
	"github.com/edvina16/icpal-backend/internal/models"
)

type TaskDB interface {
	CreateTask(ctx context.Context, params database.CreateTaskParams) (database.Task, error)
	ListTasks(ctx context.Context) ([]database.Task, error)
}

type TaskService struct {
	DB TaskDB
}

func (s *TaskService) ListTasks(ctx context.Context) ([]models.Task, error) {
	dbTasks, err := s.DB.ListTasks(ctx)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	for _, dbTask := range dbTasks {
		tasks = append(tasks, models.Task{
			ID:      int(dbTask.ID),
			Title:   dbTask.Title,
			Content: dbTask.Content,
		})
	}
	return tasks, nil
}

func (s *TaskService) CreateTask(ctx context.Context, title, content string) (models.Task, error) {
	params := database.CreateTaskParams{
		Title:   title,
		Content: content,
	}
	dbTask, err := s.DB.CreateTask(ctx, params)
	if err != nil {
		return models.Task{}, err
	}
	return models.Task{
		ID:      int(dbTask.ID),
		Title:   dbTask.Title,
		Content: dbTask.Content,
	}, nil
}
