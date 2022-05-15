-- name: CreateTask :one
INSERT INTO task (
  task_id,
  list_id,
  description,
  done,
  create_at,
  update_at
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetTask :one
SELECT * FROM task
WHERE task_id = $1 LIMIT 1;

-- name: UpdateTask :one
UPDATE task
SET list_id = $2, description = $3, done = $4, update_at = $5 
WHERE task_id = $1
RETURNING *;

-- name: ListTasks :many
SELECT * FROM task
ORDER BY task_id
LIMIT $1
OFFSET $2;

-- name: DeleteTask :exec
DELETE FROM task
WHERE task_id = $1;