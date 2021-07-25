-- name: GetGroup :one
SELECT * FROM groups
WHERE id = $1 LIMIT 1;

-- name: GetGroups :many
SELECT * FROM groups;

-- name: CreateGroup :one
INSERT INTO groups (
  organizer,
  name,
  description
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateGroup :one
UPDATE groups
SET
  organizer = $2,
  name = $3,
  description = $4
WHERE id = $1
RETURNING *;

-- name: DeleteGroup :exec
DELETE FROM groups
WHERE id = $1;