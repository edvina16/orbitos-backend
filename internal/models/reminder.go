package models

import "time"

type Reminder struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	TaskID    int       `json:"task_id"`
	Message   string    `json:"message"`
	RemindAt  time.Time `json:"remind_at"`
	Frequency string    `json:"frequency"`
	CreatedAt time.Time `json:"created_at"`
}
