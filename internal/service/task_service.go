package service

import (
	"context"

	"github.com/edvina16/atmon-backend/internal/database"
	"github.com/edvina16/atmon-backend/internal/models"
)

type TaskDB interface {
	CreateTask(ctx context.Context, params database.CreateTaskParams) (database.Task, error)
	ListTasks(ctx context.Context) ([]database.Task, error)
	DeleteTask(ctx context.Context, taskID int32) error
	UpdateTask(ctx context.Context, arg database.UpdateTaskParams) error
	GetTaskByID(ctx context.Context, taskID int32) (database.Task, error)
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

func (s *TaskService) DeleteTask(ctx context.Context, taskID int) error {
	return s.DB.DeleteTask(ctx, int32(taskID))
}

func (s *TaskService) UpdateTask(ctx context.Context, taskID int, title, content string) error {
	arg := database.UpdateTaskParams{
		ID:      int32(taskID),
		Title:   title,
		Content: content,
	}

	return s.DB.UpdateTask(ctx, arg)
}

func (s *TaskService) GetTaskByID(ctx context.Context, taskID int32) (models.Task, error) {
	dbTask, err := s.DB.ListTasks(ctx)
	if err != nil {
		return models.Task{}, err
	}
	for _, t := range dbTask {
		if t.ID == taskID {
			return models.Task{
				ID:      int(t.ID),
				Title:   t.Title,
				Content: t.Content,
			}, nil
		}
	}
	return models.Task{}, nil
}
