package service

import (
	"context"
	"testing"

	"github.com/edvina16/orbitos-backend/internal/database"
)

type mockStateDB struct{}

func (m *mockStateDB) ListStates(ctx context.Context, boardID int32) ([]database.State, error) {
	return []database.State{
		{ID: 1, Name: "TestState", BoardID: boardID},
	}, nil
}

func TestListStates(t *testing.T) {
	svc := StateService{DB: &mockStateDB{}}
	states, err := svc.ListStates(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(states) != 1 || states[0].Name != "TestState" || states[0].BoardID != 1 {
		t.Fatalf("expected one state named 'TestState' with BoardID 1, got %+v", states)
	}
}
