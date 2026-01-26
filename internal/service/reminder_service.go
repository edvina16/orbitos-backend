package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/edvina16/orbitos-backend/internal/database"
)

type ReminderDB interface {
	CreateReminder(ctx context.Context, params database.CreateReminderParams) (database.Reminder, error)
}

type ReminderService struct {
	DB ReminderDB
}

func (s *ReminderService) CreateReminder(ctx context.Context, userID int, taskID int, message string, remindAt string, frequency string) (database.Reminder, error) {
	remindAtTime, err := time.Parse(time.RFC3339, remindAt)
	if err != nil {
		return database.Reminder{}, err
	}
	params := database.CreateReminderParams{
		UserID:    int32(userID),
		TaskID:    int32(taskID),
		Message:   sql.NullString{String: message, Valid: message != ""},
		RemindAt:  remindAtTime,
		Frequency: sql.NullString{String: frequency, Valid: frequency != ""},
	}
	return s.DB.CreateReminder(ctx, params)
}
