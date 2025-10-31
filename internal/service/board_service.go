package service

import (
	"context"

	"github.com/edvina16/protai-backend/internal/database"
	"github.com/edvina16/protai-backend/internal/models"
)

type BoardDB interface {
	ListBoards(ctx context.Context) ([]database.Board, error)
	CreateBoard(ctx context.Context, name string) (database.Board, error)
	GetBoardByID(ctx context.Context, boardID int32) (database.Board, error)
	DeleteBoard(ctx context.Context, id int32) error
	UpdateBoard(ctx context.Context, arg database.UpdateBoardParams) error
}

type BoardService struct {
	DB BoardDB
}

func (s *BoardService) ListBoards(ctx context.Context) ([]models.Board, error) {
	dbBoards, err := s.DB.ListBoards(ctx)
	if err != nil {
		return nil, err
	}
	var boards []models.Board
	for _, b := range dbBoards {
		boards = append(boards, models.Board{ID: int(b.ID), Name: b.Name})
	}
	return boards, nil
}

func (s *BoardService) GetBoardByID(ctx context.Context, id int) (models.Board, error) {
	dbBoards, err := s.DB.ListBoards(ctx)
	if err != nil {
		return models.Board{}, err
	}
	for _, b := range dbBoards {
		if int(b.ID) == id {
			return models.Board{ID: int(b.ID), Name: b.Name}, nil
		}
	}
	return models.Board{}, nil
}

func (s *BoardService) CreateBoard(ctx context.Context, name string) (models.Board, error) {
	b, err := s.DB.CreateBoard(ctx, name)
	if err != nil {
		return models.Board{}, err
	}
	return models.Board{ID: int(b.ID), Name: b.Name}, nil
}

func (s *BoardService) DeleteBoard(ctx context.Context, id int) error {
	return s.DB.DeleteBoard(ctx, int32(id))
}

func (s *BoardService) UpdateBoard(ctx context.Context, id int, name string) error {
	params := database.UpdateBoardParams{
		ID:   int32(id),
		Name: name,
	}
	return s.DB.UpdateBoard(ctx, params)
}
