package service

import (
	"context"

	"github.com/edvina16/orbitos-backend/internal/database"
	"github.com/edvina16/orbitos-backend/internal/models"
)

type BoardDB interface {
	ListBoards(ctx context.Context, userID int32) ([]database.Board, error)
	CreateBoard(ctx context.Context, arg database.CreateBoardParams) (database.Board, error)
	GetBoardByID(ctx context.Context, arg database.GetBoardByIDParams) (database.Board, error)
	DeleteBoard(ctx context.Context, arg database.DeleteBoardParams) error
	UpdateBoard(ctx context.Context, arg database.UpdateBoardParams) error
}

type BoardService struct {
	DB BoardDB
}

func (s *BoardService) ListBoards(ctx context.Context, userID int) ([]models.Board, error) {
	dbBoards, err := s.DB.ListBoards(ctx, int32(userID))
	if err != nil {
		return nil, err
	}
	var boards []models.Board
	for _, b := range dbBoards {
		boards = append(boards, models.Board{ID: int(b.ID), Name: b.Name, UserID: int(b.UserID)})
	}
	return boards, nil
}

func (s *BoardService) GetBoardByID(ctx context.Context, id int, userID int) (models.Board, error) {
	params := database.GetBoardByIDParams{ID: int32(id), UserID: int32(userID)}
	dbBoard, err := s.DB.GetBoardByID(ctx, params)
	if err != nil {
		return models.Board{}, err
	}
	return models.Board{ID: int(dbBoard.ID), Name: dbBoard.Name, UserID: int(dbBoard.UserID)}, nil
}

func (s *BoardService) CreateBoard(ctx context.Context, name string, userID int) (models.Board, error) {
	params := database.CreateBoardParams{Name: name, UserID: int32(userID)}
	b, err := s.DB.CreateBoard(ctx, params)
	if err != nil {
		return models.Board{}, err
	}
	return models.Board{ID: int(b.ID), Name: b.Name, UserID: int(b.UserID)}, nil
}

func (s *BoardService) DeleteBoard(ctx context.Context, id int, userID int) error {
	params := database.DeleteBoardParams{ID: int32(id), UserID: int32(userID)}
	return s.DB.DeleteBoard(ctx, params)
}

func (s *BoardService) UpdateBoard(ctx context.Context, id int, name string, userID int) error {
	params := database.UpdateBoardParams{ID: int32(id), Name: name, UserID: int32(userID)}
	return s.DB.UpdateBoard(ctx, params)
}
