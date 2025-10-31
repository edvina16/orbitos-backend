package service

import (
	"context"

	"github.com/edvina16/icpal-backend/internal/database"
	"github.com/edvina16/icpal-backend/internal/models"
)

type BoardDB interface {
	ListBoards(ctx context.Context) ([]database.Board, error)
	CreateBoard(ctx context.Context, name string) (database.Board, error)
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
