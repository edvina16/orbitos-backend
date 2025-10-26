-- name: ListBoards :many
SELECT * FROM boards;

-- name: ListStatesByBoard :many
SELECT * FROM states WHERE board_id = $1;

-- name: ListTasksByState :many
SELECT * FROM tasks WHERE state_id = $1;

-- name: CreateBoard :one
INSERT INTO boards (name) VALUES ($1) RETURNING *;

-- name: CreateState :one
INSERT INTO states (name, board_id) VALUES ($1, $2) RETURNING *;

-- name: CreateTask :one
INSERT INTO tasks (title, content, state_id) VALUES ($1, $2, $3) RETURNING *;

-- name: ListTasks :many
SELECT * FROM tasks;