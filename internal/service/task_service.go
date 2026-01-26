package service

import (
	"context"

	"github.com/edvina16/orbitos-backend/internal/database"
	"github.com/edvina16/orbitos-backend/internal/models"
)

type TaskDB interface {
	CreateTask(ctx context.Context, params database.CreateTaskParams) (database.Task, error)
	ListTasks(ctx context.Context, userID int32) ([]database.Task, error)
	DeleteTask(ctx context.Context, arg database.DeleteTaskParams) error
	UpdateTask(ctx context.Context, arg database.UpdateTaskParams) error
	GetTaskByID(ctx context.Context, arg database.GetTaskByIDParams) (database.Task, error)
}

type TaskService struct {
	DB TaskDB
}

func (s *TaskService) ListTasks(ctx context.Context, userID int) ([]models.Task, error) {
	dbTasks, err := s.DB.ListTasks(ctx, int32(userID))
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	for _, dbTask := range dbTasks {
		tasks = append(tasks, models.Task{
			ID:      int(dbTask.ID),
			Title:   dbTask.Title,
			Content: dbTask.Content,
			UserID:  int(dbTask.UserID),
		})
	}
	return tasks, nil
}

func (s *TaskService) CreateTask(ctx context.Context, title, content string, userID int) (models.Task, error) {
	params := database.CreateTaskParams{
		Title:   title,
		Content: content,
		UserID:  int32(userID),
	}
	dbTask, err := s.DB.CreateTask(ctx, params)
	if err != nil {
		return models.Task{}, err
	}
	return models.Task{
		ID:      int(dbTask.ID),
		Title:   dbTask.Title,
		Content: dbTask.Content,
		UserID:  int(dbTask.UserID),
	}, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, taskID int, userID int) error {
	params := database.DeleteTaskParams{ID: int32(taskID), UserID: int32(userID)}
	return s.DB.DeleteTask(ctx, params)
}

func (s *TaskService) UpdateTask(ctx context.Context, taskID int, title, content string, userID int) error {
	arg := database.UpdateTaskParams{
		ID:      int32(taskID),
		Title:   title,
		Content: content,
		UserID:  int32(userID),
	}
	return s.DB.UpdateTask(ctx, arg)
}

func (s *TaskService) GetTaskByID(ctx context.Context, taskID int, userID int) (models.Task, error) {
	params := database.GetTaskByIDParams{ID: int32(taskID), UserID: int32(userID)}
	dbTask, err := s.DB.GetTaskByID(ctx, params)
	if err != nil {
		return models.Task{}, err
	}
	return models.Task{
		ID:      int(dbTask.ID),
		Title:   dbTask.Title,
		Content: dbTask.Content,
		UserID:  int(dbTask.UserID),
		StateID: int(dbTask.StateID),
	}, nil
}

func (s *TaskService) CreateTaskInState(ctx context.Context, title, content string, stateID, userID int) (models.Task, error) {
	params := database.CreateTaskParams{
		Title:   title,
		Content: content,
		StateID: int32(stateID),
		UserID:  int32(userID),
	}
	dbTask, err := s.DB.CreateTask(ctx, params)
	if err != nil {
		return models.Task{}, err
	}
	return models.Task{
		ID:      int(dbTask.ID),
		Title:   dbTask.Title,
		Content: dbTask.Content,
		UserID:  int(dbTask.UserID),
		StateID: int(dbTask.StateID),
	}, nil
}
