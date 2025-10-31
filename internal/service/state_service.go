package service

import (
	"context"

	"github.com/edvina16/icpal-backend/internal/database"
	"github.com/edvina16/icpal-backend/internal/models"
)

type StateDB interface {
	ListStatesByBoard(ctx context.Context, boardID int32) ([]database.State, error)
	CreateState(ctx context.Context, params database.CreateStateParams) (database.State, error)
}

type StateService struct {
	DB StateDB
}

func (s *StateService) ListStatesByBoard(ctx context.Context, boardID int) ([]models.State, error) {
	dbStates, err := s.DB.ListStatesByBoard(ctx, int32(boardID))
	if err != nil {
		return nil, err
	}
	var states []models.State
	for _, dbState := range dbStates {
		states = append(states, models.State{
			ID:      int(dbState.ID),
			Name:    dbState.Name,
			BoardID: int(dbState.BoardID),
		})
	}
	return states, nil
}

func (s *StateService) CreateState(ctx context.Context, name string, boardID int) (models.State, error) {
	params := database.CreateStateParams{
		Name:    name,
		BoardID: int32(boardID),
	}
	b, err := s.DB.CreateState(ctx, params)
	if err != nil {
		return models.State{}, err
	}
	return models.State{ID: int(b.ID), Name: b.Name, BoardID: int(b.BoardID)}, nil
}
