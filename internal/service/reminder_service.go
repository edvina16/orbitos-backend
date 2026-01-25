package service

import (
	"context"
	"github.com/edvina16/orbitos-backend/internal/database"
	"time"
)

type ReminderDB interface {
	CreateTask(ctx context.Context, params database.CreateReminderParams) (database.Reminder, error)
}

type ReminderService struct {
	db ReminderDB
}

func (s *ReminderService) CreateReminder(ctx context.Context, userID int, taskID int, message string, remindAt string, frequency string) (database.Reminder, error) {
	remindAtTime, err := time.Parse(time.RFC3339, remindAt)
	if err != nil {
		return database.Reminder{}, err
	}
	params := database.CreateReminderParams{
		UserID:    int32(userID),
		TaskID:    int32(taskID),
		Message:   message,
		RemindAt:  remindAtTime,
		Frequency: frequency,
	}
	return s.db.CreateTask(ctx, params)
}
