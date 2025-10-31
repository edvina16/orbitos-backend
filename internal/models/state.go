package models

type State struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	BoardID int    `json:"board_id"`
	UserID  int    `json:"user_id"`
}
