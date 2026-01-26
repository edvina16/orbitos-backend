package service

import (
	"context"

	"github.com/edvina16/orbitos-backend/internal/database"
	"github.com/edvina16/orbitos-backend/internal/models"
)

type StateDB interface {
	ListStatesByBoard(ctx context.Context, arg database.ListStatesByBoardParams) ([]database.State, error)
	CreateState(ctx context.Context, arg database.CreateStateParams) (database.State, error)
	ListTasksByState(ctx context.Context, arg database.ListTasksByStateParams) ([]database.Task, error)
	CreateTask(ctx context.Context, arg database.CreateTaskParams) (database.Task, error)
	UpdateTaskState(ctx context.Context, arg database.UpdateTaskStateParams) error
	DeleteState(ctx context.Context, arg database.DeleteStateParams) error
	UpdateState(ctx context.Context, arg database.UpdateStateParams) error
	GetStateByID(ctx context.Context, arg database.GetStateByIDParams) (database.State, error)
}

type StateService struct {
	DB StateDB
}

func (s *StateService) ListStatesByBoard(ctx context.Context, boardID int, userID int) ([]models.State, error) {
	params := database.ListStatesByBoardParams{BoardID: int32(boardID), UserID: int32(userID)}
	dbStates, err := s.DB.ListStatesByBoard(ctx, params)
	if err != nil {
		return nil, err
	}
	var states []models.State
	for _, dbState := range dbStates {
		states = append(states, models.State{
			ID:      int(dbState.ID),
			Name:    dbState.Name,
			BoardID: int(dbState.BoardID),
			UserID:  int(dbState.UserID),
		})
	}
	return states, nil
}

func (s *StateService) CreateState(ctx context.Context, name string, boardID int, userID int) (models.State, error) {
	params := database.CreateStateParams{Name: name, BoardID: int32(boardID), UserID: int32(userID)}
	b, err := s.DB.CreateState(ctx, params)
	if err != nil {
		return models.State{}, err
	}
	return models.State{ID: int(b.ID), Name: b.Name, BoardID: int(b.BoardID), UserID: int(b.UserID)}, nil
}

func (s *StateService) ListTasksByState(ctx context.Context, stateID int, userID int) ([]models.Task, error) {
	params := database.ListTasksByStateParams{StateID: int32(stateID), UserID: int32(userID)}
	dbTasks, err := s.DB.ListTasksByState(ctx, params)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	for _, dbTask := range dbTasks {
		tasks = append(tasks, models.Task{
			ID:      int(dbTask.ID),
			Title:   dbTask.Title,
			Content: dbTask.Content,
			StateID: int(dbTask.StateID),
			UserID:  int(dbTask.UserID),
		})
	}
	return tasks, nil
}

func (s *StateService) CreateTaskInState(ctx context.Context, title string, content string, stateID int) (models.Task, error) {
	params := database.CreateTaskParams{
		Title:   title,
		Content: content,
		StateID: int32(stateID),
	}
	t, err := s.DB.CreateTask(ctx, params)
	if err != nil {
		return models.Task{}, err
	}
	return models.Task{ID: int(t.ID), Title: t.Title, Content: t.Content, StateID: int(t.StateID)}, nil
}

func (s *StateService) UpdateTaskState(ctx context.Context, taskID int, stateID int, userID int) error {
	params := database.UpdateTaskStateParams{
		ID:      int32(taskID),
		StateID: int32(stateID),
		UserID:  int32(userID),
	}
	return s.DB.UpdateTaskState(ctx, params)
}

func (s *StateService) DeleteState(ctx context.Context, stateID int, userID int) error {
	params := database.DeleteStateParams{ID: int32(stateID), UserID: int32(userID)}
	return s.DB.DeleteState(ctx, params)
}

func (s *StateService) UpdateState(ctx context.Context, stateID int, name string) error {
	params := database.UpdateStateParams{
		ID:   int32(stateID),
		Name: name,
	}
	return s.DB.UpdateState(ctx, params)
}

func (s *StateService) GetStateByID(ctx context.Context, stateID int32, userID int32) (models.State, error) {
	params := database.GetStateByIDParams{ID: stateID, UserID: userID}
	dbState, err := s.DB.GetStateByID(ctx, params)
	if err != nil {
		return models.State{}, err
	}
	return models.State{
		ID:      int(dbState.ID),
		Name:    dbState.Name,
		BoardID: int(dbState.BoardID),
		UserID:  int(dbState.UserID),
	}, nil
}
