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

-- name: UpdateTaskState :exec
UPDATE tasks SET state_id = $2 WHERE id = $1;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;

-- name: UpdateTask :exec
UPDATE tasks SET title = $2, content = $3 WHERE id = $1;

-- name: DeleteState :exec
DELETE FROM states WHERE id = $1;

-- name: UpdateState :exec
UPDATE states SET name = $2 WHERE id = $1;

-- name: DeleteBoard :exec
DELETE FROM boards WHERE id = $1;

-- name: UpdateBoard :exec
UPDATE boards SET name = $2 WHERE id = $1;

-- name: GetBoardByID :one
SELECT * FROM boards WHERE id = $1;

-- name: GetStateByID :one
SELECT * FROM states WHERE id = $1;

-- name: GetTaskByID :one
SELECT * FROM tasks WHERE id = $1;