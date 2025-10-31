-- name: ListBoards :many
SELECT * FROM boards WHERE user_id = $1;

-- name: ListStatesByBoard :many
SELECT * FROM states WHERE board_id = $1 AND user_id = $2;

-- name: ListTasksByState :many
SELECT * FROM tasks WHERE state_id = $1 AND user_id = $2;

-- name: CreateBoard :one
INSERT INTO boards (name, user_id) VALUES ($1, $2) RETURNING *;

-- name: CreateState :one
INSERT INTO states (name, board_id, user_id) VALUES ($1, $2, $3) RETURNING *;

-- name: CreateTask :one
INSERT INTO tasks (title, content, state_id, user_id) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: ListTasks :many
SELECT * FROM tasks WHERE user_id = $1;

-- name: UpdateTaskState :exec
UPDATE tasks SET state_id = $2 WHERE id = $1 AND user_id = $3;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1 AND user_id = $2;

-- name: UpdateTask :exec
UPDATE tasks SET title = $2, content = $3 WHERE id = $1 AND user_id = $4;

-- name: DeleteState :exec
DELETE FROM states WHERE id = $1 AND user_id = $2;

-- name: UpdateState :exec
UPDATE states SET name = $2 WHERE id = $1 AND user_id = $3;

-- name: DeleteBoard :exec
DELETE FROM boards WHERE id = $1 AND user_id = $2;

-- name: UpdateBoard :exec
UPDATE boards SET name = $2 WHERE id = $1 AND user_id = $3;

-- name: GetBoardByID :one
SELECT * FROM boards WHERE id = $1 AND user_id = $2;

-- name: GetStateByID :one
SELECT * FROM states WHERE id = $1 AND user_id = $2;

-- name: GetTaskByID :one
SELECT * FROM tasks WHERE id = $1 AND user_id = $2;

-- name: ListBoardsByUser :many
SELECT * FROM boards WHERE user_id = $1;

-- name: ListTasksByUser :many
SELECT * FROM tasks WHERE user_id = $1;

-- name: CreateUser :one
INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;
