-- name: ListTasks :many
SELECT id, title, content FROM tasks ORDER BY id DESC;

-- name: CreateTask :one
INSERT INTO tasks (title, content) VALUES ($1, $2)
RETURNING id, title, content;