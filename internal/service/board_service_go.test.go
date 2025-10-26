package service

import (
	"context"
	"testing"

	"github.com/edvina16/icpal-backend/internal/database"
)

type mockBoardDb struct{}

func (m *mockBoardDb) ListBoards(ctx context.Context) ([]database.Board, error) {
	return []database.Board{
		{ID: 1, Name: "Test Board"},
	}, nil
}

func (m *mockBoardDb) GetBoardByID(ctx context.Context, id int32) (database.Board, error) {
	return database.Board{ID: id, Name: "Test Board"}, nil
}

func (m *mockBoardDb) CreateBoard(ctx context.Context, name string) (database.Board, error) {
	return database.Board{ID: 2, Name: name}, nil
}

func TestListBoards(t *testing.T) {
	svc := BoardService{DB: &mockBoardDb{}}
	boards, err := svc.ListBoards(context.Background())
	if err != nil {
		t.Fatalf("unexpected error listing boards: %v", err)
	}
	if len(boards) != 1 || boards[0].Name != "Test Board" {
		t.Fatalf("expected one board with the name 'Test Board', got %v", boards)
	}
}

func TestGetBoardByID(t *testing.T) {
	svc := BoardService{DB: &mockBoardDb{}}
	board, err := svc.GetBoardByID(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error getting board by ID: %v", err)
	}
	if board.ID != 1 || board.Name != "Test Board" {
		t.Fatalf("expected board with ID 1 and name 'Test Board', got %+v", board)
	}
}

func TestCreateBoard(t *testing.T) {
	svc := BoardService{DB: &mockBoardDb{}}
	board, err := svc.CreateBoard(context.Background(), "New Board")
	if err != nil {
		t.Fatalf("unexpected error creating board: %v", err)
	}
	if board.Name != "New Board" {
		t.Fatalf("unexpected board: %+v", board)
	}
}
