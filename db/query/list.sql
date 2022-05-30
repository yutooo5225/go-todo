-- name: CreateList :one
INSERT INTO list (
  list_id,
  list_name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetList :one
SELECT * FROM list
WHERE list_id = $1 LIMIT 1;

-- name: UpdateList :one
UPDATE list
SET list_name = $2
WHERE list_id = $1
RETURNING *;

-- name: Lists :many
SELECT * FROM list
ORDER BY list_id
LIMIT $1
OFFSET $2;

-- name: DeleteList :exec
DELETE FROM list
WHERE list_id = $1;